package workerpool

func newWorker(id int) worker {
	return &workman{
		id: id,
	}
}

type worker interface {
	ProcessTask(task Task) error
	ID() int
}

type workman struct {
	id int
}

func (w *workman) ProcessTask(task Task) error {
	err := task.Exec()
	return err
}

func (w *workman) ID() int {
	return w.id
}
