package cache

import (
	"github.com/robinjoseph08/redisqueue/v2"

	"github.com/mss-boot-io/mss-boot/service/store/pkg/storage"
)

type Message struct {
	redisqueue.Message
}

func (m *Message) GetID() string {
	return m.ID
}

func (m *Message) GetStream() string {
	return m.Stream
}

func (m *Message) GetValues() map[string]interface{} {
	return m.Values
}

func (m *Message) SetID(id string) {
	m.ID = id
}

func (m *Message) SetStream(stream string) {
	m.Stream = stream
}

func (m *Message) SetValues(values map[string]interface{}) {
	m.Values = values
}

func (m *Message) GetPrefix() (prefix string) {
	if m.Values == nil {
		return
	}
	v, _ := m.Values[storage.PrefixKey]
	prefix, _ = v.(string)
	return
}

func (m *Message) SetPrefix(prefix string) {
	if m.Values == nil {
		m.Values = make(map[string]interface{})
	}
	m.Values[storage.PrefixKey] = prefix
}
