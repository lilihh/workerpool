# v.1.0.0

#### Created Time: 2020-08-05

## Main Function:
* Start()
* Close()
* ReceiveNormalTask(task Task) error
* ReceiveUrgentTask(task Task) error

## Assist Function:
* Debug()

## Architecture
```text
    workerpool
     __________________________________________________________
    |
    |   dispatcher                          workers
    |   contain 2 storage                   implement by goruntine
    |   implement by channel                grab and process task
    |                                                          _
    |   ------------------                                      |
    |   urgent tasks        (grab first)    |---->  worker #1   |
    |   ██ ██ ██ ██ ██ ██  -----------------|                   |
    |   ------------------                  |---->  worker #2   |
    |                                       |                   |
    |                                       |---->  worker #3    \
    |   ------------------                  |                     workerNum
    |   normal tasks        (grab later)    |---->  worker #4    / 
    |   ██ ██ ██ ██ ██ ██  - - - - - - - - -|                   | 
    |   ------------------                  |---->  worker #5   |
    |                                       .                   |
    |                                       .                   |      
    |   |<---  buf   --->|                  .                  _|
    |
```
