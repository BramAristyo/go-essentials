package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker Pool
// A group of workers sitting idle, waiting for work to come in.
// When a job arrives, one worker picks it up, processes it, and goes back to waiting.
// You control how many workers exist, not how many jobs come in.

// func worker(jobs <-chan Job, results chan<- Result) {
// 	for {
// 		select {
// 		case job, ok := <-jobs:
// 			if !ok {
// 				return
// 			}
//
// 			results <- Result{JobID: job.ID, Output: fmt.Sprintf("Hello from %s", job.ID)}
// 		}
// 	}
// }

func main() {
	const TOTAL_WORKER_PROCESS = 5

	jobs := make(chan Job, TOTAL_WORKER_PROCESS)
	results := make(chan Result, TOTAL_WORKER_PROCESS)

	// init worker based on TOTAL_WORKER_PROCESS
	for range TOTAL_WORKER_PROCESS {
		go Worker(jobs, results)
	}

	// send a job to the worker
	for j := 1; j <= TOTAL_WORKER_PROCESS; j++ {
		jobs <- Job{ID: j}
	}

	close(jobs) // close jobs channel after sending all jobs

	// get result from channel results
	var res Result
	for range TOTAL_WORKER_PROCESS {
		res = <-results
		fmt.Println(res.Output)
	}

	close(results)
}

type Job struct {
	ID int
}

type Result struct {
	JobID  int
	Output string
}

func Worker(jobs <-chan Job, results chan<- Result) {
	// pakai for range untuk menunggu channel jobs di send
	for job := range jobs {
		time.Sleep(time.Second)
		results <- Result{JobID: job.ID, Output: fmt.Sprintf("Hello from %d", job.ID)}
	}
}

func WorkerWithWG(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Second)
		results <- Result{JobID: job.ID, Output: fmt.Sprintf("Hello from %d", job.ID)}
	}
}
