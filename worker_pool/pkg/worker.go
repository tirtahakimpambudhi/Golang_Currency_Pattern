package pkg

import (
	"fmt"
	"sync"
)

type Task struct {
	ID int
}

func (t *Task) Run() {
	fmt.Println(t.ID)
}

type WorkerPool struct {
	Tasks    []Task
	Worker   int
	TaskChan chan Task
	wg       sync.WaitGroup
}

func (w *WorkerPool) work() {
	for task := range w.TaskChan {
		task.Run()
	}
	w.wg.Done()
}

func (w *WorkerPool) Run() {
	w.TaskChan = make(chan Task, len(w.Tasks))

	w.wg.Add(len(w.Tasks))
	defer w.wg.Done()
	for i := 0; i < w.Worker; i++ {
		go w.work()
	}
	for _, task := range w.Tasks {
		w.TaskChan <- task
	}
	close(w.TaskChan)
}
