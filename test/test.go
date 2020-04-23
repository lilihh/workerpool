package main

import (
	"sync"
	"workerpool"
)

var wp workerpool.IWorkerPool

func main() {
	// initialize
	wp = workerpool.NewWorkerPool(5, 2)
	wp.Debug(true)

	wp.Start()
	defer wp.Close()

	// run
	wg := &sync.WaitGroup{}
	tasks := make([]workerpool.Task, 0, 10)
	for i := 0; i < 10; i++ {
		tasks = append(tasks, newTestTask(wg, i+1))
	}

	wg.Add(len(tasks))
	for _, task := range tasks {
		// 全部做完
		for err := wp.ReceiveTask(task); err != nil; {
			err = wp.ReceiveTask(task)
		}
	}

	wg.Wait()
}
