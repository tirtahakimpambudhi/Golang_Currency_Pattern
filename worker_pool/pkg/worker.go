package pkg

import (
	"go_wp/contract"
	"sync"
)


type WorkerPool struct {
	Tasks    []contract.Task
	Worker   int
	TaskChan chan contract.Task
	wg       sync.WaitGroup
}

func (w *WorkerPool) work() {
	for task := range w.TaskChan {
		task.Proccess()
		w.wg.Done()
	}
}

func (w *WorkerPool) Run() {
	w.TaskChan = make(chan contract.Task, len(w.Tasks))

	w.wg.Add(len(w.Tasks))
	defer w.wg.Wait()
	//receieve work 
	for i := 0; i < w.Worker; i++ {
		go w.work()
	}
	//send work
	for _, task := range w.Tasks {
		w.TaskChan <- task
	}
	close(w.TaskChan)
}

