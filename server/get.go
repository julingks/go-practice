package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
}
