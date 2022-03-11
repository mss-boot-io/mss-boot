// Package memory is a memory source
package memory

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mss-boot-io/mss-boot/core/config/source"
)

type memory struct {
	sync.RWMutex
	ChangeSet *source.ChangeSet
	Watchers  map[string]*watcher
}

// Read read
func (s *memory) Read() (*source.ChangeSet, error) {
	s.RLock()
	cs := &source.ChangeSet{
		Format:    s.ChangeSet.Format,
		Timestamp: s.ChangeSet.Timestamp,
		Data:      s.ChangeSet.Data,
		Checksum:  s.ChangeSet.Checksum,
		Source:    s.ChangeSet.Source,
	}
	s.RUnlock()
	return cs, nil
}

// Watch watch
func (s *memory) Watch() (source.Watcher, error) {
	w := &watcher{
		id:      uuid.New().String(),
		updates: make(chan *source.ChangeSet, 100),
		source:  s,
	}

	s.Lock()
	s.Watchers[w.id] = w
	s.Unlock()
	return w, nil
}

// Write write
func (s *memory) Write(cs *source.ChangeSet) error {
	s.Update(cs)
	return nil
}

// Update allows manual updates of the cfg data.
func (s *memory) Update(c *source.ChangeSet) {
	// don't process nil
	if c == nil {
		return
	}

	// hash the file
	s.Lock()
	// update changeset
	s.ChangeSet = &source.ChangeSet{
		Data:      c.Data,
		Format:    c.Format,
		Source:    "memory",
		Timestamp: time.Now(),
	}
	s.ChangeSet.Checksum = s.ChangeSet.Sum()

	// update watchers
	for _, w := range s.Watchers {
		select {
		case w.updates <- s.ChangeSet:
		default:
		}
	}
	s.Unlock()
}

// String string
func (s *memory) String() string {
	return "memory"
}

// NewSource new a memory source
func NewSource(opts ...source.Option) source.Source {
	var options source.Options
	for _, o := range opts {
		o(&options)
	}

	s := &memory{
		Watchers: make(map[string]*watcher),
	}

	if options.Context != nil {
		c, ok := options.Context.Value(changeSetKey{}).(*source.ChangeSet)
		if ok {
			s.Update(c)
		}
	}

	return s
}
