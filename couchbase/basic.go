package main

import (
	"fmt"
	"log"

	"github.com/couchbaselabs/go-couchbase"
)

func mf(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v", msg, err)
	}
}

func main() {

	url := "http://localhost:8091"

	c, err := couchbase.Connect(url)
	mf(err, "connect - "+url)

	p, err := c.GetPool("default")
	mf(err, "pool")

	b, err := p.GetBucketWithAuth("bucketname", "bucketname", "pw")

	mf(err, "bucket")

	err = b.Set("test", 90, map[string]interface{}{"x": 1})
	mf(err, "set")

	ob := map[string]interface{}{}
	err = b.Get("test", &ob)
	mf(err, "get")

	fmt.Println("Done.")
}
