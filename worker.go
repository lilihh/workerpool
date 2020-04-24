package workerpool

import (
	"log"
	"time"
)

func newWorker(id int, dispatcher *dispatcher) *worker {
	return &worker{
		id:             id,
		quit:           make(chan bool),
		taskDepositary: dispatcher,
	}
}

type worker struct {
	id             int         // 工人編號
	isLog          bool        // 是否開啟紀錄
	quit           chan bool   // 關閉執行緒
	taskDepositary *dispatcher // 工作倉庫
}

func (w *worker) start() {
	go func() {
		for {
			var task Task

			select {
			case <-w.quit:
				return
			case task = <-w.taskDepositary.priorityTasks:
				// grab priority task fisrt
			default:
				select {
				case task = <-w.taskDepositary.normalTasks:
					// grab normal task later
				default:
				}
			}

			if task != nil {
				err := w.processTask(task)

				if w.isLog {
					if err != nil {
						log.Printf("worker #%d is done with error:%v\n", w.id, err.Error())
					} else {
						log.Printf("worker #%d is done with no error\n", w.id)
					}
				}
			}

			time.Sleep(time.Millisecond) // release CPU utilization
		}
	}()
}

func (w *worker) close() {
	w.quit <- true
}

func (w *worker) processTask(task Task) error {
	err := task.Exec()
	return err
}

func (w *worker) log(ok bool) {
	w.isLog = ok
}
