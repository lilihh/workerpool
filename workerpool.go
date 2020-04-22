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
	ReceiveTasks(tasks ...Task)

	IsLog(ok bool)
	AllowTaskOverFlow(ok bool)
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

func (wp *workerPool) ReceiveTasks(tasks ...Task) {
	if wp.allowTaskOverFlow {
		for _, task := range tasks {
			wp.dispatcher.receiveTaskEvenFlood(task)
		}

	} else {
		for _, task := range tasks {
			if err := wp.dispatcher.receiveTask(task); err != nil {
				return
			}
		}
	}
}

func (wp *workerPool) IsLog(ok bool) {
	wp.dispatcher.isLog(ok)
	wp.pool.isLog(ok)
}

func (wp *workerPool) AllowTaskOverFlow(ok bool) {
	wp.allowTaskOverFlow = ok
}

// TODO: set worker waiting time
// func (wp *workerPool) SetWaitingTime(dur *time.Duration) {

// }

// TODO: define error
// TODO: optimize log
// TODO: write testing
// TODO: write README What is workerpool?
