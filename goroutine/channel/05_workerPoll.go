package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job..")
		time.Sleep(3 * time.Second)
		results <- j * (j + 1)
		fmt.Println("worker ", id, " finished job")
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	fmt.Println("Success init 5 worker, for this step the worker wait the jobs comes in")
	time.Sleep(2 * time.Second)

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("success init job ", j)
		time.Sleep(time.Second)
	}

	fmt.Println("Success give a worker jobs to process the 'j' value")

	for a := 1; a <= 5; a++ {
		fmt.Println("Success get Result: ", <-results)
	}
}
