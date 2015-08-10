package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	b, _ := ioutil.ReadFile("pack.zip")

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		panic(err)
	}

	for _, zf := range r.File {
		fmt.Println(zf.Name)

		dst, err := os.Create("dst.zip")
		if err != nil {
			panic(err)
		}

		defer dst.Close()

		src, err := zf.Open()
		if err != nil {
			panic(err)
		}
		defer src.Close()
		io.Copy(dst, src)
	}
}
