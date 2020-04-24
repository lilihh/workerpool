package workerpool

import "fmt"

// NewWorkerPool return a pool with workers
func NewWorkerPool(buf, numOfWorkers int) IWorkerPool {
	if buf < 0 || numOfWorkers < 1 {
		return nil
	}

	wp := &workerPool{
		dispatcher: newDispatcher(buf),
		workers:    make([]*worker, 0, numOfWorkers),
	}

	for i := 0; i < numOfWorkers; i++ {
		wp.workers = append(wp.workers, newWorker(i+1, wp.dispatcher))
	}

	return wp
}

// IWorkerPool will receive tasks and dispatch them to workers
type IWorkerPool interface {
	Start()
	Close()
	ReceiveTask(task Task, isPriority bool) error

	Debug(ok bool)
}

type workerPool struct {
	dispatcher *dispatcher
	workers    []*worker

	allowTaskOverFlow bool
}

func (wp *workerPool) Start() {
	for _, worker := range wp.workers {
		worker.start()
	}
}

func (wp *workerPool) Close() {
	for _, worker := range wp.workers {
		worker.close()
	}
}

func (wp *workerPool) ReceiveTask(task Task, isPriority bool) error {
	if task == nil {
		return fmt.Errorf("illeagalArgument")
	}

	if isPriority {
		if err := wp.dispatcher.receivePriorityTask(task); err != nil {
			return err
		}
	} else {
		if err := wp.dispatcher.receiveNormalTask(task); err != nil {
			return err
		}
	}

	return nil
}

func (wp *workerPool) Debug(ok bool) {
	for _, worker := range wp.workers {
		worker.log(ok)
	}
}

// TODO: define error
// TODO: optimize log
// TODO: write testing
