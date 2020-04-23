package main

import (
	"fmt"
	"sync"
	"time"
	"workerpool"
)

func newTestTask(wg *sync.WaitGroup, i int) workerpool.Task {
	return &testTask{
		wg: wg,
		id: i,
	}
}

type testTask struct {
	wg *sync.WaitGroup
	id int
}

func (t *testTask) Exec() error {
	defer func() {
		t.wg.Done()
	}()

	fmt.Println(t.id)
	time.Sleep(time.Millisecond)
	return nil
}
