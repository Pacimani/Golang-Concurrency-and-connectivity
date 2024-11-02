package main


import (
	"fmt"
	"io"
	"os"
)

// capper implements io.WriteCloser and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {

	// convert everything to uppercase
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, c := range p {

		if c >= 'a' && c <= 'z' {
			c -= diff
			//continue
		}
		out[i] = c
	}

	return c.wtr.Write(out)
}

func main() {

	c := Capper{os.Stdout}
	fmt.Fprint(&c, "Hello there\n")	
}