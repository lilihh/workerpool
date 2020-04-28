package main

import (
	"sync"
	"workerpool"
)

var wp workerpool.IWorkerPool

func main() {
	// initialize
	wp = workerpool.NewWorkerPool(3, 2)
	wp.Debug(true)

	wp.Start()
	defer wp.Close()

	// run
	wg := &sync.WaitGroup{}
	tasks := make([]workerpool.Task, 0, 15)
	for i := 0; i < 15; i++ {
		tasks = append(tasks, newTestTask(wg, i+1))
	}

	wg.Add(len(tasks))
	// test normal or prioriy? choose one please!

	// // normal test
	// for _, task := range tasks {

	// 	for err := wp.ReceiveTask(task, false); err != nil; {
	// 		err = wp.ReceiveTask(task, false)
	// 	}
	// }

	// test priority
	for i, task := range tasks {

		if (i/5)%2 == 0 {
			for err := wp.ReceiveTask(task, false); err != nil; {
				err = wp.ReceiveTask(task, false)
			}
		} else {
			for err := wp.ReceiveTask(task, true); err != nil; {
				err = wp.ReceiveTask(task, true)
			}
		}
	}

	wg.Wait()
}
