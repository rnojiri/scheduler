package scheduler

import (
	"sync"
	"time"
)

// manager - schedules all expression executions
type manager struct {
	taskMap sync.Map
}

// Job - a job to be executed
type job func()

// Task - a scheduled task
type Task struct {
	ID       string
	Duration time.Duration
	Job      job
	running  uint32
}

// Manager - the scheduler manager interface
type Manager interface {
	// AddTask - adds a new task
	AddTask(task *Task, autoStart bool) error

	// Exists - checks if a task exists
	Exists(id string) bool

	// IsRunning - checks if a task is running
	IsRunning(id string) bool

	// RemoveTask - removes a task
	RemoveTask(id string) bool

	// RemoveAllTasks - removes all tasks
	RemoveAllTasks()

	// StopTask - stops a task
	StopTask(id string) error

	// StartTask - starts a task
	StartTask(id string) error

	// GetNumTasks - returns the number of tasks
	GetNumTasks() int

	// GetTasksIDs - returns a list of task IDs
	GetTasksIDs() []string

	// GetTasks - returns a list of tasks
	GetTasks() []*Task

	// GetTask - returns a task by it's ID
	GetTask(id string) *Task
}

var _ (Manager) = (*manager)(nil)
