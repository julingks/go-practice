package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

/*
func (f *rot13Reader) Read(p []byte) (n int, err error) {
	return 0, nil
}

*/
func main() {
	s := strings.NewReader("Lbh penqxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
