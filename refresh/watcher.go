package refresh

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	*fsnotify.Watcher
	*Manager
	context context.Context
}

func NewWatcher(r *Manager) *Watcher {
	w, _ := fsnotify.NewWatcher()

	return &Watcher{
		Watcher: w,
		Manager: r,
		context: r.context,
	}

}

func (w *Watcher) Start() {
	go func() {
		for {
			err := filepath.Walk(w.AppRoot, func(path string, info os.FileInfo, err error) error {
				if info == nil {
					w.cancelFunc()
					return errors.New("nil directory!")
				}
				if info.IsDir() {
					if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") || w.isIgnoredFolder(path) {
						return filepath.SkipDir
					}
				}
				if w.isWatchedFile(path) {
					w.Add(path)
				}
				return nil
			})

			if err != nil {
				w.context.Done()
				break
			}
			// sweep for new files every 1 second
			time.Sleep(1 * time.Second)
		}
	}()
}

func (w Watcher) isIgnoredFolder(path string) bool {
	for _, e := range w.IgnoredFolders {
		rel, err := filepath.Rel(e, path)
		if err != nil {
			// unable to construct relative path, not an ignored folder
			continue
		}

		if strings.Contains(rel, "..") {
			// to construct a relative path requires going up the directory tree, not
			// an ignored folder
			continue
		}

		return true
	}
	return false
}

func (w Watcher) isWatchedFile(path string) bool {
	ext := filepath.Ext(path)

	for _, e := range w.IncludedExtensions {
		if strings.TrimSpace(e) == ext {
			return true
		}
	}

	return false
}
