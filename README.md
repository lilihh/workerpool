# workerpool

## What is workerpool?
Workerpool is a pool containing several workers, who are waiting to process task send by you.

## Implement
* Every worker is a goroutine/thread, and start working when `Start()` is be called.
* There's a channel of `Task` in dispatcher, storing tasks send by `ReceiveTask(task Task)`

* Structure Diagram
```go
    workerpool
     __________________________________________________________
    |
    |                           goruntine
    |                           grab and process task
    |                                              _
    |                                               |
    |   dispatcher              |---->  worker #1   |
    |   ------------------      |                   |
    |   tasks                   |---->  worker #2   |
    |   ██ ██ ██ ██ ██ ██   ----|                    \
    |   ------------------      |---->  worker #3     workerNum
    |                           |                    /
    |   |<--   buf   -->|       |---->  worker #4   |
    |                           .                   |
    |                           .                   |      
    |                           .                  _|
    |
```

## Installation
    $ go get github.com/lilihh/workerpool

## How to use?
This section will show some examples.

### Simplest example
Let's begin with the simplest one.

```go
// define your own task
func newExampleTask() workerpool.Task {
    // generate a task and return it
}

type exampleTask struct {}

func (t *exampleTask) Exec() error {
    // do something
}

func main() {
    // new a workerpool
    wp := workerpool.NewWorkerPool(buf, workerNum)
    wp.Start()
    defer wp.Close()

    // generate tasks
    num := 10
    tasks := make([]workerpool.Task, 0, num)
    for i := 0; i < num; i++ {
        tasks = append(tasks, newExampleTask())
    }

    // perocess tasks
    for _, task := range tasks {
        wp.ReceiveTask(task)
    }

    // wait
    <-time.After(time.Millisecond)  
}
```

### Example with every task must be done
The capacity of task-buffer in the workerpool is constant. What about the amount of tasks is larger than the capacity?
`ReceiveTask(task Task)` will return an error if the workerpool receive a task but the buffer is full already. In that case, workerpool do not receive that task actually. So you have to control it by yourself.

```go
// if you want every task must be done anyway
func main() {
    // new a workerpool
    ...

    // generate tasks
    ...

    // perocess tasks
    for _, task := range tasks {
        for err := wp.ReceiveTask(task); err != nil; {
            err = wp.ReceiveTask(task)
        }
    }

    // wait
    ...
}

```

### Example with WaitGroup
In real case, we usually use [`sync.WaitGroup`](https://golang.org/pkg/sync/#WaitGroup) instead of `time.After` to control the process.

```go
func newExampleTask(wg *sync.WaitGroup){} Task {
    return &exampleTask{
        wg: wg,
    }
}

type exampleTask struct {
    wg *sync.WaitGroup
}

func (t *exampleTask) Exec() error {
    // do something

    t.wg.Done()
}
```

```go
func main() {
    // new a workerpool
    ...

    // generate tasks
    wg := &sync.WaitGroup{}

	tasks := make([]workerpool.Task, 0, 10)
	for i := 0; i < 10; i++ {
		tasks = append(tasks, newExampleTask(wg, fmt.Sprintf("%s",i+1)))
	}

    // perocess tasks
	wg.Add(len(tasks))
	
    for _, task := range tasks {
        for err := wp.ReceiveTask(task); err != nil; {
            err = wp.ReceiveTask(task)
        }
    }

    wg.Wait()
}
```

### Option
if you want to know what happen in workerpool, you can use `Debug()`
```go
// new a workerpool
...
wp.Debug(true)
wp.Start()
```