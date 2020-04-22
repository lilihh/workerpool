package main

import (
	"time"
	"workerpool"
)

var wp workerpool.IWorkerPool
var service *server

func main() {
	// initialize
	wp = workerpool.NewWorkerPool(5, 2)
	wp.AllowTaskOverFlow(false)
	wp.IsLog(true)

	wp.Start()
	defer wp.Close()

	// run
	service := &server{}
	go service.run()

	// just for waiting
	<-time.After(5 * time.Second)
}

type server struct{}

func (server) run() {
	tasks := make([]workerpool.Task, 0, 10)
	for i := 0; i < 10; i++ {
		tasks = append(tasks, newTestTask(i+1))
	}

	wp.ReceiveTasks(tasks...)
}
