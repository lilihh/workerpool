package workerpool

import (
	"log"
	"sync"
)

func newPool(numOfWorker int, waitGroup *sync.WaitGroup) iPool {
	p := &pool{
		wg:          waitGroup,
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
	wg          *sync.WaitGroup
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
					log.Println("worker #", worker.ID(), " is done with error:", err)
				} else {
					log.Println("worker #", worker.ID(), " is done with no error")
				}
			}

			p.wg.Done()
		}()
	}
}
