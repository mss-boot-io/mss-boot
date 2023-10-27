package task

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/2/21 15:35:43
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/2/21 15:35:43
 */

import (
	"context"
	"log/slog"

	"github.com/robfig/cron/v3"
)

var task = &Server{
	opts: setDefaultOption(),
}

// Server manage
type Server struct {
	ctx  context.Context
	opts options
}

// New server
func New(opts ...Option) *Server {
	task.Options(opts...)
	return task
}

// GetJob get job
func GetJob(key string) (string, cron.Job, bool) {
	task.opts.mux.Lock()
	defer task.opts.mux.Unlock()
	s, ok := task.opts.schedules[key]
	if !ok {
		return "", nil, false
	}
	return s.spec, s.job, true
}

// UpdateJob update or create job
func UpdateJob(key string, spec string, job cron.Job) error {
	task.opts.mux.Lock()
	defer task.opts.mux.Unlock()
	s, ok := task.opts.schedules[key]
	if ok {
		task.opts.task.Remove(s.entryID)
	}
	entryID, err := task.opts.task.AddJob(spec, job)
	if err != nil {
		slog.Error("task add job error", slog.Any("err", err))
		return err
	}
	task.opts.schedules[key] = schedule{
		spec:    spec,
		job:     job,
		entryID: entryID,
	}
	return nil
}

// RemoveJob remove job
func RemoveJob(key string) error {
	task.opts.mux.Lock()
	defer task.opts.mux.Unlock()
	s, ok := task.opts.schedules[key]
	if !ok {
		return nil
	}
	task.opts.task.Remove(s.entryID)
	delete(task.opts.schedules, key)
	return nil
}

// Options set options
func (e *Server) Options(opts ...Option) {
	for _, o := range opts {
		o(&e.opts)
	}
}

// String server name
func (e *Server) String() string {
	return "task"
}

// Start server
func (e *Server) Start(ctx context.Context) error {
	var err error
	e.ctx = ctx
	for i, s := range e.opts.schedules {
		s.entryID, err = e.opts.task.AddJob(e.opts.schedules[i].spec, e.opts.schedules[i].job)
		if err != nil {
			slog.ErrorContext(ctx, "task add job error", slog.Any("err", err))
			return err
		}
		e.opts.schedules[i] = s
	}
	go func() {
		e.opts.task.Run()
		<-ctx.Done()
		err = e.Shutdown(ctx)
		if err != nil {
			slog.ErrorContext(ctx, e.String()+" Server shutdown error", slog.Any("err", err.Error()))
		}
	}()
	return nil
}

// Shutdown server
func (e *Server) Shutdown(_ context.Context) error {
	e.opts.task.Stop()
	return nil
}
