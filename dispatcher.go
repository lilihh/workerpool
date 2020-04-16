package workerpool

import (
	"log"
	"sync"
)

func newDispatcher(buf int, waitGroup *sync.WaitGroup) iDispatcher {
	return &dispatcher{
		wg:          waitGroup,
		taskStorage: make(chan Task, buf),
		quit:        make(chan bool),
	}
}

type iDispatcher interface {
	start(pool iPool)
	close()
	isLog(ok bool)
	receiveTask(task Task)
}

type dispatcher struct {
	wg          *sync.WaitGroup
	taskStorage chan Task
	quit        chan bool

	islog bool
}

func (d *dispatcher) start(pool iPool) {
	if d.islog {
		log.Println("dispatcher start")
	}

	go d.launchTask(pool)
}

func (d *dispatcher) close() {
	d.quit <- true
}

func (d *dispatcher) isLog(ok bool) {
	d.islog = ok
}

func (d *dispatcher) receiveTask(task Task) {
	d.taskStorage <- task
}

func (d *dispatcher) launchTask(pool iPool) {
	if d.islog {
		defer log.Println("dispatcher close")
	}

	for {
		select {
		case task := <-d.taskStorage:
			pool.launchTaskToWorker(task)
		case <-d.quit:
			return
		}
	}
}
