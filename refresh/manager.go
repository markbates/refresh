package refresh

import (
	"os/exec"
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

func New(c *Configuration) *Manager {
	return &Manager{
		Configuration: c,
		Logger:        NewLogger(c),
		Restart:       make(chan bool),
		gil:           &sync.Once{},
	}
}

func (r *Manager) Start() error {
	w := NewWatcher(r)
	w.Start()
	go r.build(fsnotify.Event{Name: ":start:"})
	go func() {
		for {

			event := <-w.Events
			if event.Op != fsnotify.Chmod {
				go r.build(event)
			}
			w.Remove(event.Name)
			w.Add(event.Name)
		}
	}()
	go func() {
		for {
			err := <-w.Errors
			r.Logger.Error(err)
		}
	}()
	r.runner()
	return nil
}

func (r *Manager) build(event fsnotify.Event) {
	r.gil.Do(func() {
		time.Sleep(r.BuildDelay * time.Millisecond)

		now := time.Now()
		r.Logger.Print("Rebuild on: %s", event.Name)
		cmd := exec.Command("go", "build", "-v", "-i", "-o", r.FullBuildPath())
		err := r.runAndListen(cmd, func(s string) {
			r.Logger.Print(s)
		})

		if err != nil {
			r.Logger.Error("Building Error!")
			r.Logger.Error(err)
		} else {
			tt := time.Since(now)
			r.Logger.Success("Building Completed (PID: %d) (Time: %s)", cmd.Process.Pid, tt)
			r.Restart <- true
		}
		r.gil = &sync.Once{}
	})
}
