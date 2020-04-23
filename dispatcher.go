package workerpool

import (
	"fmt"
	"log"
)

func newDispatcher(buf int) iDispatcher {
	return &dispatcher{
		taskStorage: make(chan Task, buf),
		quit:        make(chan bool),
	}
}

type iDispatcher interface {
	start(pool iPool)
	close()
	receiveTask(task Task) error

	isLog(ok bool)
}

type dispatcher struct {
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

func (d *dispatcher) receiveTask(task Task) error {
	select {
	case d.taskStorage <- task:
		return nil
	default:
		return fmt.Errorf("buffer over flow")
	}
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
