/*
 * @Author: lwnmengjing
 * @Date: 2022/3/21 14:51
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/21 14:51
 */

package cfg

import (
	"errors"
	"github.com/mss-boot-io/mss-boot/service/store/pkg/storage"
	"github.com/mss-boot-io/mss-boot/service/store/pkg/storage/cache"
	"github.com/mss-boot-io/mss-boot/service/store/pkg/storage/queue"
)

type Memory struct {
	PoolSize uint `yaml:"poolSize" json:"poolSize"`
}

// String string
func (Memory) String() string {
	return "memory"
}

func (e *Memory) GetCache() (storage.AdapterCache, error) {
	return cache.NewMemory(), nil
}

func (e *Memory) GetQueue() (storage.AdapterQueue, error) {
	return queue.NewMemory(e.PoolSize), nil
}

func (e *Memory) GetLocker() (storage.AdapterLocker, error) {
	return nil, errors.New("method GetLocker not implemented")
}
