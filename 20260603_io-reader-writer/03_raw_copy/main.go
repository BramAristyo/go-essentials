package main

import (
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	src := strings.NewReader("This data flows from Reader to Writer")

	buf := make([]byte, 4)
	// n, err := src.Read(buf)
	// if err != nil && err != io.EOF {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// os.Stdout.Write(buf[:n])

	for {
		n, err := src.Read(buf)

		time.Sleep(1 * time.Second)

		if n > 0 {
			os.Stdout.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}
	}
}
