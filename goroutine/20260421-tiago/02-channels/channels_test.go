package main

import (
	"fmt"
	"testing"
	"time"
)

/*
   go test -v channels_test.go
*/

func TestBasicChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		fmt.Println("cooking start")
		time.Sleep(2 * time.Second)
		ch <- "🍕🍕🍕"
		fmt.Println("done!")
	}()

	fmt.Println("wait pizza .... ")
	pizza := <-ch
	fmt.Println("nom nom ... ", pizza)

	if pizza != "🍕🍕🍕" {
		t.Errorf("Expected 🍕🍕🍕, but got %s", pizza)
	}
}

func TestSelect(t *testing.T) {
	chPizza := make(chan string)
	chMatcha := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		chPizza <- "🍕"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chMatcha <- "🧃"
	}()

	select {
	case serve := <-chPizza:
		fmt.Println("received ", serve)
	case serve := <-chMatcha:
		fmt.Println("received ", serve)
	case <-time.After(5 * time.Second):
		t.Fatal("timeout: service too slow!")
	}
}

func TestForeverSelect(t *testing.T) {
	chPizza := make(chan string)
	chMatcha := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chPizza <- "🍕"
		time.Sleep(1 * time.Second)
		chPizza <- "🍕"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		chMatcha <- "🧃"
	}()

	count := 0
Loop:
	for {
		select {
		case serve := <-chPizza:
			fmt.Println("received ", serve)
			count++
		case serve := <-chMatcha:
			fmt.Println("received ", serve)
			count++
		case <-time.After(3 * time.Second):
			fmt.Println("No more signals, stopping...")
			break Loop
		}
	}

	fmt.Printf("Total items received: %d\n", count)
	if count == 0 {
		t.Error("Should have received at least one item")
	}
}
