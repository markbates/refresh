package refresh

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Manager struct {
	*Configuration
	Logger  *log.Logger
	Restart chan bool
	gil     *sync.Once
}

func (r *Manager) build() {
	r.gil.Do(func() {
		sl := time.Duration(r.BuildDelay) * time.Millisecond
		time.Sleep(sl)

		b := NewBuilder(*r)
		err := b.Build()
		if err != nil {
			r.Logger.Printf("Error building: %s, %s", b.ID, err)
		}

		r.Restart <- true
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
			// fmt.Printf("### event -> %s\n", event)
			// log.Println("modified file:", event.Name)
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
		Logger:        log.New(os.Stdout, "refresh: ", log.LstdFlags),
		Restart:       make(chan bool),
		gil:           &sync.Once{},
	}
}
