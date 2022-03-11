package env

import (
	"github.com/mss-boot-io/mss-boot/core/config/source"
)

type watcher struct {
	exit chan struct{}
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	<-w.exit

	return nil, source.ErrWatcherStopped
}

func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func newWatcher() (source.Watcher, error) {
	return &watcher{exit: make(chan struct{})}, nil
}
