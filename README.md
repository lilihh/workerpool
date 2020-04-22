# workerpool

## What is workerpool?

## Installation
    $ go get github.com/lilihh/workerpool

## Example
```go
// define your own task
type exampleTask struct {}

func (exampleTask) Exec() error {
    // do something
}

func main() {
    // new a workerpool
    wp := workerpool.NewWorkerPool(buf, workerNum)
    wp.Start()
    defer wp.Close()

    // **optional**
    // wp.IsLog(true)

    // push tasks into the workerpool
    tasks := []exampleTask{}{...}
    wp.ReceiveTasks(tasks...)
}
```