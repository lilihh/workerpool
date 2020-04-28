package workerpool

func newDispatcher(buf int) *dispatcher {
	return &dispatcher{
		normalTasks: make(chan Task, buf),
		urgentTasks: make(chan Task, buf),
	}
}

type dispatcher struct {
	normalTasks chan Task // storage for normal tasks
	urgentTasks chan Task // storage for urgent tasks
}

func (d *dispatcher) receiveNormalTask(task Task) error {
	select {
	case d.normalTasks <- task:
		return nil
	default:
		return newError(NormalBufferOverflow)
	}
}

func (d *dispatcher) receiveUrgentTask(task Task) error {
	select {
	case d.urgentTasks <- task:
		return nil
	default:
		return newError(UrgentBufferOverflow)
	}
}
