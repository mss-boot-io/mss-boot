/*
 * @Author: lwnmengjing
 * @Date: 2022/3/22 9:25
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/22 9:25
 */

package cfg

import (
	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/service/store/pkg/storage"
	"strings"
)

// Provider provider
type Provider struct {
	Memory *Memory `yaml:"memory" json:"memory"`
	Redis  *Redis  `yaml:"redis" json:"redis"`
	NSQ    *NSQ    `yaml:"nsq" json:"nsq"`
}

func (e Provider) Init(cache, queue, locker string) {
	if (cache == "memory" || queue == "memory") &&
		e.Memory == nil {
		log.Fatal("please config memory provider")
	}
	if (cache == "redis" || queue == "redis" || locker == "redis") &&
		e.Redis == nil {
		log.Fatal("please config redis provider")
	}
	if queue == "nsq" && e.NSQ == nil {
		log.Fatal("please config nsq provider")
	}
	e.initCache(cache)
	e.initQueue(queue)
	e.initLocker(locker)

}

func (e *Provider) initCache(provider string) {
	var err error
	switch strings.ToLower(provider) {
	case "memory":
		storage.Cache, err = e.Memory.GetCache()
	case "redis":
		storage.Cache, err = e.Redis.GetCache()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (e *Provider) initQueue(provider string) {
	var err error
	switch strings.ToLower(provider) {
	case "memory":
		storage.Queue, err = e.Memory.GetQueue()
	case "redis":
		storage.Queue, err = e.Redis.GetQueue()
	case "nsq":
		storage.Queue, err = e.NSQ.GetQueue()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (e *Provider) initLocker(provider string) {
	var err error
	switch strings.ToLower(provider) {
	case "redis":
		storage.Locker, err = e.Redis.GetLocker()
	}
	if err != nil {
		log.Fatal(err)
	}
}
