package test

import (
	"fmt"
	"go_wp/contract"
	"go_wp/pkg"
	"testing"
	"time"
)

const totalTask int = 100
const totalWorker int = 20 

type Task struct {
	ID int
}

func (t *Task) Proccess() {
	fmt.Println(t.ID)
	time.Sleep(2*time.Second)
}


func TestTaskWorkerPool(t *testing.T) {
	var tasks []contract.Task
	for i := 1; i <= totalTask; i++ {
		tasks = append(tasks, &Task{ID: i})
	}

	wp := pkg.WorkerPool{Tasks: tasks, Worker: totalWorker}
	wp.Run()
}
