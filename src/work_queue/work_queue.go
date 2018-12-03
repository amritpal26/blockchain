package work_queue

import (
	// "fmt"
)

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
	StopRequest chan int
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
	// initialize struct; start nWorkers workers as goroutines
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, maxJobs)
	q.StopRequest = make(chan int, maxJobs)

	for i := uint(0); i < nWorkers; i++{
		go q.worker()
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
	// Listen on the .Jobs channel for incoming tasks. For each task...
	// run tasks by calling .Run(),
	// send the return value back on Results channel.
	// Exit (return) when .Jobs is closed.
func (queue WorkQueue) worker() {
	run := true
	for run{
		// fmt.Println("listening for new jobs")
		job := <- queue.Jobs
		// fmt.Println("Got a new jobs")
		queue.Results <- job.Run() 
		if len(queue.StopRequest) > 0{
			run = false
			return
		} 
	}
}

// put the work into the Jobs channel so a worker can find it and start the task.
func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
}

// close .Jobs and remove all remaining jobs from the channel.
func (queue WorkQueue) Shutdown() {
	queue.StopRequest <- 1 
}
