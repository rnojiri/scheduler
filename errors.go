package scheduler

import "errors"

var (
	ErrTaskAlreadyExists  error = errors.New("task id already exists")
	ErrTaskAlreadyRunning error = errors.New("task id already running")
	ErrTaskAlreadyStopped error = errors.New("task id already stopped")
	ErrTaskNotExists      error = errors.New("task id not exists")
)
