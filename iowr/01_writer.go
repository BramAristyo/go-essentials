package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func writeToWriter(w io.Writer, message string) {
	w.Write([]byte(message))
}

func main() {
	// os.Stdout implement io.Writer, so we can pass it direct as a Writer
	writeToWriter(os.Stdout, "Hello from Writer")

	var sb strings.Builder
	writeToWriter(&sb, "this data goes into memory")

	fmt.Println("sb value is: ", sb.String())
}
