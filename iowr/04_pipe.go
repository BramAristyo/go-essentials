package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Println(w, "chunk 1")
		time.Sleep(1 * time.Second)
		time.Sleep(1 * time.Second)
		time.Sleep(1 * time.Second)
		fmt.Println(w, "chunk 2")
		time.Sleep(1 * time.Second)
		fmt.Println(w, "chunk 3")
		time.Sleep(1 * time.Second)
		w.Close()
	}()

	io.Copy(os.Stdout, r)
}
