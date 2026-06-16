package main

import (
	"fmt"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	ID    int
	Value int
}

func worker2(jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		results <- Result{ID: j.ID, Value: j.Value + 1}
	}
}

func main() {
	const workerCount = 5
	jobs := make(chan Job, workerCount)
	results := make(chan Result, workerCount)

	for range workerCount {
		go worker2(jobs, results)
	}

	for wc := range workerCount {
		jobs <- Job{wc, wc}
	}

	close(jobs)

	mergeResult := Result{100, 0}
	var temporaryRes Result
	for range workerCount {
		temporaryRes = <-results
		mergeResult.Value += temporaryRes.Value
	}

	fmt.Println(mergeResult.Value, " <- this is the final value merged")

}
