package memory

import (
	"github.com/mss-boot-io/mss-boot/core/config/source"
)

type watcher struct {
	id      string
	updates chan *source.ChangeSet
	source  *memory
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	cs := <-w.updates
	return cs, nil
}

func (w *watcher) Stop() error {
	w.source.Lock()
	delete(w.source.Watchers, w.id)
	w.source.Unlock()
	return nil
}
