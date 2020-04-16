package workerpool

// Task should be implement by user
type Task interface {
	Exec() error
}
