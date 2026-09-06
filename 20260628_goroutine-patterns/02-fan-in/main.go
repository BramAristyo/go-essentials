package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.IntN(3)) * time.Second)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	// its a fan in pattern, N different sender channel to ONE receiver channel
	c := fanIn(boring("Joe"), boring("Ann"))

	for range 6 {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring!")
}
