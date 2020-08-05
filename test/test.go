package main

import (
	"sync"
	"workerpool"
)

var wp workerpool.IWorkerPool

func main() {
	// initialize
	wp = workerpool.NewWorkerPool(3, 2)
	wp.Debug()

	wp.Start()
	defer wp.Close()

	// run
	wg := &sync.WaitGroup{}
	tasks := make([]workerpool.Task, 0, 15)
	for i := 0; i < 15; i++ {
		tasks = append(tasks, newTestTask(wg, i+1))
	}

	wg.Add(len(tasks))

	// test priority
	for i, task := range tasks {

		if (i/5)%2 == 0 {
			for err := wp.ReceiveNormalTask(task); err != nil; {
				err = wp.ReceiveNormalTask(task)
			}
		} else {
			for err := wp.ReceiveUrgentTask(task); err != nil; {
				err = wp.ReceiveUrgentTask(task)
			}
		}
	}

	wg.Wait()
}
