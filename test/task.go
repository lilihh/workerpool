package main

import (
	"fmt"
	"time"
	"workerpool"
)

func newTestTask(i int) workerpool.Task {
	return &testTask{
		id: i,
	}
}

type testTask struct {
	id int
}

func (t *testTask) Exec() error {
	fmt.Println(t.id)
	time.Sleep(time.Millisecond)
	return nil
}
