package refresh

import (
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Manager struct {
	*Configuration
	Logger  *Logger
	Restart chan bool
	gil     *sync.Once
}

func (r *Manager) build() {
	r.gil.Do(func() {
		time.Sleep(r.BuildDelay * time.Millisecond)

		b := NewBuilder(*r)
		err := b.Build()
		if err == nil {
			r.Restart <- true
		}
		r.gil = &sync.Once{}
	})
}

func (r *Manager) Start() error {
	w := NewWatcher(r)
	w.Start()
	go r.build()
	go func() {
		for {

			event := <-w.Events
			if event.Op != fsnotify.Chmod {
				go r.build()
			}
			w.Remove(event.Name)
			w.Add(event.Name)
		}
	}()
	go func() {
		for {
			err := <-w.Errors
			log.Println("error:", err)
		}
	}()
	r.runner()
	return nil
}

func New(c *Configuration) *Manager {
	return &Manager{
		Configuration: c,
		Logger:        NewLogger(c),
		Restart:       make(chan bool),
		gil:           &sync.Once{},
	}
}
