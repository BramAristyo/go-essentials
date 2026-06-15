package main

import (
	"fmt"
	"runtime"
)

func sayHelloTo(name string, message chan string) {
	var data = fmt.Sprintf("hello %s", name)

	message <- data
}

func main()  {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	go sayHelloTo("wick")
}
