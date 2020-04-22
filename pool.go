package workerpool

import (
	"log"
)

func newPool(numOfWorker int) iPool {
	p := &pool{
		idleWorkers: make(chan worker, numOfWorker),
	}

	for i := 0; i < numOfWorker; i++ {
		p.idleWorkers <- newWorker(i + 1)
	}

	return p
}

type iPool interface {
	isLog(ok bool)
	launchTaskToWorker(task Task)
}

type pool struct {
	idleWorkers chan worker

	islog bool
}

func (p *pool) isLog(ok bool) {
	p.islog = ok
}

func (p *pool) launchTaskToWorker(task Task) {
	select {
	case worker := <-p.idleWorkers:
		go func() {
			defer func() { p.idleWorkers <- worker }()

			err := worker.ProcessTask(task)
			if p.islog {
				if err != nil {
					log.Printf("worker #%d is done with error:%v\n", worker.ID(), err.Error())
				} else {
					log.Printf("worker #%d is done with no error\n", worker.ID())
				}
			}
		}()
	}
}
