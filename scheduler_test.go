package scheduler_test

import (
	"testing"
	"time"

	"github.com/rnojiri/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestTaskAutoStarting(t *testing.T) {

	count := 0

	task := scheduler.NewTask("foo", 10*time.Millisecond,
		func() {
			count++
		},
	)

	m := scheduler.New()
	err := m.AddTask(task, true)
	assert.NoError(t, err, "expected no error adding task")

	assert.Equal(t, 1, m.GetNumTasks(), "expects only one task stored")
	assert.True(t, m.IsRunning("foo"), "expected the task to be running")

	<-time.After(11 * time.Millisecond)
	assert.Equal(t, 1, count, "expected the count to be: 1")

	<-time.After(5 * time.Millisecond)
	assert.Equal(t, 1, count, "expected the count to not be incremented")

	<-time.After(10 * time.Millisecond)
	assert.Equal(t, 2, count, "expected the count to be: 2")
}

func TestTaskStartingManually(t *testing.T) {

	count := 0

	task := scheduler.NewTask("foo", 10*time.Millisecond,
		func() {
			count++
		},
	)

	m := scheduler.New()
	err := m.AddTask(task, false)
	assert.NoError(t, err, "expected no error adding task")

	assert.Equal(t, 1, m.GetNumTasks(), "expects only one task stored")
	assert.False(t, m.IsRunning("foo"), "expected the task to not be running")

	<-time.After(11 * time.Millisecond)
	assert.Zero(t, count, "expected the count to not be incremented")

	err = m.StartTask("foo")
	assert.NoError(t, err, "expected no error starting task")

	<-time.After(11 * time.Millisecond)
	assert.Equal(t, 1, count, "expected the count to be: 1")
}

func TestStopTask(t *testing.T) {

	count := 0

	task := scheduler.NewTask("foo", 10*time.Millisecond,
		func() {
			count++
		},
	)

	m := scheduler.New()
	err := m.AddTask(task, true)
	assert.NoError(t, err, "expected no error adding task")

	assert.Equal(t, 1, m.GetNumTasks(), "expects only one task stored")
	assert.True(t, m.IsRunning("foo"), "expected the task to be running")

	<-time.After(11 * time.Millisecond)
	assert.Equal(t, 1, count, "expected the count to be: 1")

	err = m.StopTask("foo")
	assert.NoError(t, err, "expected no error stopping task")

	<-time.After(20 * time.Millisecond)
	assert.Equal(t, 1, count, "expected the count to stay equals: 1")

	err = m.StartTask("foo")
	assert.NoError(t, err, "expected no error starting task")

	<-time.After(11 * time.Millisecond)
	assert.Equal(t, 2, count, "expected the count to be: 2")
}

func TestGetTaskByName(t *testing.T) {

	task := scheduler.NewTask("foo", 10*time.Millisecond, func() {})

	m := scheduler.New()
	err := m.AddTask(task, true)
	assert.NoError(t, err, "expected no error adding task")

	assert.True(t, m.Exists("foo"), "expects the task to exist")
	assert.Equal(t, task, m.GetTask("foo"), "expects the same task")
}

func TestGetTasks(t *testing.T) {

	tasks := []*scheduler.Task{
		scheduler.NewTask("foo1", 10*time.Millisecond, func() {}),
		scheduler.NewTask("foo2", 30*time.Millisecond, func() {}),
	}

	m := scheduler.New()
	err := m.AddTask(tasks[0], true)
	assert.NoError(t, err, "expected no error adding task 1")

	err = m.AddTask(tasks[1], false)
	assert.NoError(t, err, "expected no error adding task 2")

	assert.True(t, m.Exists("foo1"), "expects the task to exist")
	assert.True(t, m.Exists("foo2"), "expects the task to exist")

	tasksStored := m.GetTasks()
	assert.ElementsMatch(t, tasks, tasksStored, "expects the same tasks stored")
}
