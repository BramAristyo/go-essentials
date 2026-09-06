package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, hasMore := <-jobs
			if hasMore {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job")
	}

	close(jobs)
	fmt.Println("hi runtime, all sender is Done their work")

	<-done // just send the done data to the nothing
	_, ok := <-jobs
	fmt.Println("received more jobs: ", ok)

}
