package testing

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(out, "Go!")
}

func main() {
	buffer := &bytes.Buffer{}
	Countdown(buffer)
}
