package workerpool

import (
	"sync"
	"time"
)

// NewWorkerPool return a pool with workers
func NewWorkerPool(buf, workers int) IWorkerPool {
	var waitGroup sync.WaitGroup
	return &workerPool{
		wg:         &waitGroup,
		dispatcher: newDispatcher(buf, &waitGroup),
		pool:       newPool(workers, &waitGroup),
	}
}

// IWorkerPool will receive tasks and dispatch them to workers
type IWorkerPool interface {
	Start()
	Close()
	ReceiveTasks(tasks []Task)

	IsLog(ok bool)
}

type workerPool struct {
	wg         *sync.WaitGroup
	dispatcher iDispatcher
	pool       iPool
}

func (wp *workerPool) Start() {
	wp.dispatcher.start(wp.pool)
}

func (wp *workerPool) Close() {
	wp.dispatcher.close()
}

func (wp *workerPool) ReceiveTasks(tasks []Task) {
	wp.wg.Add(len(tasks))

	for _, task := range tasks {
		wp.dispatcher.receiveTask(task)
	}
	wp.wg.Wait()
}

func (wp *workerPool) IsLog(ok bool) {
	wp.dispatcher.isLog(ok)
	wp.pool.isLog(ok)
}

// TODO
func (wp *workerPool) SetWaitingTime(dur *time.Duration) {

}

// TODO: optimize log
