package routinepool

import (
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

// Task .
type Task func() error

// RoutinePool .
type RoutinePool struct {
	MaxSize      int
	WorkerNum    int
	InProcessNum int32
	TaskChan     chan Task
	ErrChan      chan error
}

// NewRoutinePool .
func NewRoutinePool(size int) *RoutinePool {
	pool := &RoutinePool{
		MaxSize:      size,
		WorkerNum:    size / 2,
		InProcessNum: 0,
		TaskChan:     make(chan Task, size/2),
		ErrChan:      make(chan error, size/2),
	}
	for i := 0; i < size/2; i++ {
		go pool.background()
	}
	return pool
}

// Submit .
func (r *RoutinePool) Submit(task Task) error {
	r.submit(task, false)
	logrus.Infof("WorkerNum: %d, InProcessNum: %d", r.WorkerNum, atomic.LoadInt32(&r.InProcessNum))
	return nil
}

func (r *RoutinePool) submit(task Task, withPriority bool) error {
	if !withPriority {
		r.TaskChan <- task
		return nil
	}
	return nil
}

func (r *RoutinePool) background() {
	for {
		t := <-r.TaskChan
		atomic.AddInt32(&r.InProcessNum, 1)
		err := t()
		if err != nil {
			r.ErrChan <- err
		}
		atomic.AddInt32(&r.InProcessNum, -1)
	}
}
