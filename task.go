package scheduler

import (
	"sync/atomic"
	"time"
)

//
// A single task to run a repetitive task
// author: rnojiri
//

// NewTask - creates a new task
// job should be a func() or implements the interface scheduler.Job
func NewTask(id string, duration time.Duration, job any) *Task {

	return &Task{
		ID:       id,
		Duration: duration,
		Job:      job,
		running:  0,
	}
}

// Start - starts to run this task
func (t *Task) Start() {

	if atomic.LoadUint32(&t.running) == 1 {
		return
	}

	go func() {
		for {
			<-time.After(t.Duration)

			if atomic.LoadUint32(&t.running) == 0 {
				return
			}

			if f, casted := t.Job.(func()); casted {
				f()
				continue
			}

			if f, casted := t.Job.(Job); casted {
				f.Execute()
				continue
			}
		}
	}()

	atomic.StoreUint32(&t.running, 1)
}

// Stop - stops the task
func (t *Task) Stop() {

	atomic.StoreUint32(&t.running, 0)
}
