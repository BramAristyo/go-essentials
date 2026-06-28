package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		// the channel (sender) will be blocking until it has a receiver
		// jadi ga masalah kalau forever karena by default emang sudah ada reeivernnya
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.IntN(3)) * time.Second)
		}
	}()

	return c
}

func main() {
	c := boring("boring!")
	d := boring("d said it's boring!")

	for range 5 {
		// the issue in this case is whenever c or d done the boring first,
		// both has to block each other to continue their work
		fmt.Printf("You say %q\n", <-c)
		fmt.Printf("You say %q\n", <-d)
	}

	fmt.Println("You re boring, im leaving.")
}
