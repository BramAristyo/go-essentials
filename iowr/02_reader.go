package main

import (
	"fmt"
	"io"
	"strings"
)

func readFromReader(r io.Reader) string {
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("error: ", err)
		return ""
	}

	return string(buf[:n])
}

func main() {
	r := strings.NewReader("Hello from Reader!")

	res := readFromReader(r)
	fmt.Println("Result:", res)

	r2 := strings.NewReader("Hello from Reader2!")
	buf := make([]byte, 4)
	for {
		n, err := r2.Read(buf)

		if n > 0 {
			fmt.Printf("chunk : %q\n", buf[:n])
		}

		if err == io.EOF {
			fmt.Println("done, no more data")
			break
		}

		if err != nil {
			fmt.Println("error: ", err)
			break
		}
	}

}
