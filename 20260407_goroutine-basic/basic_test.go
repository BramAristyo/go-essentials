package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutineIsLight(t *testing.T) {
	var wg sync.WaitGroup
	counter := 100

	for i := range counter {
		wg.Add(1)
		go func(id int) {
			fmt.Println("Release Goroutine with ID: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Already release %d Goroutines", counter)
}

func goroutineLoop(wg *sync.WaitGroup, code string, counter int) {
	defer wg.Done()
	for i := range counter {
		fmt.Printf("Hi, im %s %d\n", code, i)
	}
	time.Sleep(500 * time.Millisecond)
}

func TestBasicGoroutine(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	go goroutineLoop(&wg, "A", 50)
	go goroutineLoop(&wg, "B", 50)
	wg.Wait()
}
