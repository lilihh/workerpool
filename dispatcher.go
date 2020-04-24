package workerpool

import (
	"fmt"
)

func newDispatcher(buf int) *dispatcher {
	return &dispatcher{
		normalTasks:   make(chan Task, buf),
		priorityTasks: make(chan Task, buf),
	}
}

type dispatcher struct {
	normalTasks   chan Task // 一般工作儲存庫
	priorityTasks chan Task // 優先工作儲存庫
}

func (d *dispatcher) receiveNormalTask(task Task) error {
	select {
	case d.normalTasks <- task:
		return nil
	default:
		return fmt.Errorf("buffer over flow")
	}
}

func (d *dispatcher) receivePriorityTask(task Task) error {
	select {
	case d.priorityTasks <- task:
		return nil
	default:
		return fmt.Errorf("buffer over flow")
	}
}
