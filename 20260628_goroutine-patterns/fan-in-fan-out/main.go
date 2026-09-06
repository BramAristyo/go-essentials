package main

import (
	"fmt"
	"sync"
	"time"
)

// fan-out is splitting work accross multiple goroutines to process data and receive the result on pararel (its like ./cmd/workerpool/main.go)
// fan-out is process or merged the multiple results into a final result

// https://medium.com/@sanhdoan/fan-in-and-fan-out-patterns-in-go-supercharge-your-concurrent-operations-e490e8e84994

type user struct {
	ID       int
	Name     string
	Email    string
	Location string
}

func fetchUsers() []user {
	return []user{
		{1, "Alice", "alice@example.com", "New York"},
		{2, "Bob", "bob@example.com", "San Francisco"},
		{3, "Charlie", "charlie@example.com", "Chicago"},
		{4, "Diana", "diana@example.com", "Houston"},
		{5, "Eve", "eve@example.com", "Phoenix"},
		{6, "Frank", "frank@example.com", "Philadelphia"},
		{7, "Grace", "grace@example.com", "San Antonio"},
		{8, "Hank", "hank@example.com", "San Diego"},
		{9, "Ivy", "ivy@example.com", "Dallas"},
		{10, "Jack", "jack@example.com", "Austin"},
		{11, "Karen", "karen@example.com", "Jacksonville"},
		{12, "Leo", "leo@example.com", "Seattle"},
		{13, "Mia", "mia@example.com", "Denver"},
		{14, "Nick", "nick@example.com", "Boston"},
		{15, "Olivia", "olivia@example.com", "Nashville"},
	}
}

func enrichUserData(user user) user {
	time.Sleep(1 * time.Second)
	user.Location = user.Location + ", USA"
	return user
}

func fanOut(users []user, workerCount int) []<-chan user {
	channels := make([]<-chan user, workerCount)

	for i := range workerCount {
		ch := make(chan user)
		channels[i] = ch

		go func(ch chan user, workerID int) {
			defer close(ch)
			for j := workerID; j < len(users); j += workerCount {
				enriched := enrichUserData(users[j])
				ch <- enriched
			}

		}(ch, i)
	}

	return channels
}

func fanIn(channels []<-chan user) <-chan user {
	res := make(chan user)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan user) {
			defer wg.Done()

			for user := range ch {
				res <- user
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	users := fetchUsers()
	const TOTAL_WORKER_PROCESS = 3

	start := time.Now()

	channels := fanOut(users, TOTAL_WORKER_PROCESS)
	results := fanIn(channels)

	processedUsers := []user{}
	for user := range results {
		processedUsers = append(processedUsers, user)
	}

	elapsed := time.Since(start)
	fmt.Printf("Processed %d users in %v\n", len(processedUsers), elapsed)
}
