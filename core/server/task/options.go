package task

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 16:23:53
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 16:23:53
 */

import (
	"sync"

	"github.com/robfig/cron/v3"
)

// Option params set
type Option func(*options)

type schedule struct {
	spec    string
	job     cron.Job
	entryID cron.EntryID
}

type options struct {
	task      *cron.Cron
	schedules map[string]schedule
	mux       sync.Mutex
}

// WithSchedule set schedule
func WithSchedule(key string, spec string, job cron.Job) Option {
	return func(o *options) {
		o.mux.Lock()
		o.schedules[key] = schedule{
			spec: spec,
			job:  job,
		}
		o.mux.Unlock()
	}
}

func setDefaultOption() options {
	return options{
		task:      cron.New(cron.WithSeconds(), cron.WithChain()),
		schedules: make(map[string]schedule),
	}
}
