package filenotify

// This file is borrowed from the github.com/moby/moby project.
// Licensed under the Apache License, Version 2.0

import "github.com/fsnotify/fsnotify"

// fsNotifyWatcher wraps the fsnotify package to satisfy the FileNotifier interface
type fsNotifyWatcher struct {
	*fsnotify.Watcher
}

// Events returns the fsnotify event channel receiver
func (w *fsNotifyWatcher) Events() <-chan fsnotify.Event {
	return w.Watcher.Events
}

// Errors returns the fsnotify error channel receiver
func (w *fsNotifyWatcher) Errors() <-chan error {
	return w.Watcher.Errors
}
