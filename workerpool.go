package workerpool

// NewWorkerPool return a pool with workers
func NewWorkerPool(buf, workers int) IWorkerPool {
	return &workerPool{
		dispatcher: newDispatcher(buf),
		pool:       newPool(workers),
	}
}

// IWorkerPool will receive tasks and dispatch them to workers
type IWorkerPool interface {
	Start()
	Close()
	ReceiveTask(task Task) error

	Debug(ok bool)
}

type workerPool struct {
	dispatcher iDispatcher
	pool       iPool

	allowTaskOverFlow bool
}

func (wp *workerPool) Start() {
	wp.dispatcher.start(wp.pool)
}

func (wp *workerPool) Close() {
	wp.dispatcher.close()
}

func (wp *workerPool) ReceiveTask(task Task) error {
	if err := wp.dispatcher.receiveTask(task); err != nil {
		return err
	}
	return nil
}

func (wp *workerPool) Debug(ok bool) {
	wp.dispatcher.isLog(ok)
	wp.pool.isLog(ok)
}

// TODO: set worker waiting time
// func (wp *workerPool) SetWaitingTime(dur *time.Duration) {
// }

// TODO: define error
// TODO: optimize log
// TODO: write testing
// TODO: write README What is workerpool?
