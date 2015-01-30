package main

import (
	"flag"
	"fmt"
	"github.com/couchbaselabs/go-couchbase"
	"github.com/kr/pretty"
	"log"
	"time"
)

var poolName = flag.String("pool", "default", "Pool name")
var writeFlag = flag.Bool("write", false, "If true, will write a value to the key")

func main() {
	flag.Parse()
	connectStr := "http://localhost:8091"
	bucketName := "default"
	keyName := "a"

	if len(flag.Args()) < 3 {
		//log.Fatalf("Usage: observe [-pool poolname] [-write] server bucket key")
	} else {
		connectStr = flag.Arg(0)
		bucketName = flag.Arg(1)
		keyName = flag.Arg(2)
	}

	c, err := couchbase.Connect(connectStr)
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	fmt.Printf("Connected to ver=%s\n", c.Info.ImplementationVersion)

	pool, err := c.GetPool(*poolName)
	if err != nil {
		log.Fatalf("Can't get pool %q, %v", *poolName, err)
	}

	pretty.Println("pool : ", pool)

	bucket, err := pool.GetBucket(bucketName)
	if err != nil {
		log.Fatalf("Can't get bucket %q: %v", *poolName, err)
	}
	pretty.Println("bucket : ", bucket)

	vbServerMap := bucket.VBServerMap()
	pretty.Print("vbServerMap : ", vbServerMap)

	ddocs, _ := bucket.GetDDocs()
	pretty.Println("ddocs : ", ddocs)

	key := keyName

	result, err := bucket.Observe(key)
	if err != nil {
		log.Fatalf("Observe returned error %v", err)
	}
	pretty.Printf("Observe result: %v\n", result)

	fmt.Println(*writeFlag)

	if *writeFlag {
		log.Printf("Now writing to key %q with presistence...", key)
		start := time.Now()
		err = bucket.Write(key, 0, 0, "observe test", couchbase.Persist)
		if err != nil {
			log.Fatalf("Write returned error %v", err)
		}
		log.Printf("Write with persistence took %s", time.Since(start))
	}

	fmt.Println("Done.")
}
