package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/couchbase/cbauth"
	"github.com/couchbaselabs/go-couchbase"
)

var serverURL = flag.String("serverURL", "http://localhost:8091", "couchbase server URL")
var poolName = flag.String("poolName", "default", "pool name")
var bucketName = flag.String("bucketName", "bucket", "bucket name")
var authUser = flag.String("authUser", "bucket", "auth user name (probably same as bucketName)")
var authPaswd = flag.String("authPswd", "password", "auth password")

func main() {
	flag.Parse()
	url, err := url.Parse(*serverURL)

	if err != nil {
		log.Printf("Failed to parse url %v", err)
		return
	}
	hostPort := url.Host

	user, bucket_password, err := cbauth.GetHTTPServiceAuth(hostPort)

	if err != nil {
		log.Printf("Failed %v", err)
		return
	}

	log.Printf(" HTTP Serve username %s password %s", user, bucket_password)

	client, err := couchbase.ConnectWithAuthCreds(*serverURL, user, bucket_password)

	if err != nil {
		log.Printf("Failed %v", err)
		return
	}

	cbpool, err := client.GetPool("default")
	if err != nil {
		log.Printf("Failed to connect to default pool %v", err)
		return
	}

	mUser, mPassword, err := cbauth.GetMemcachedServiceAuth(hostPort)
	if err != nil {
		log.Printf("Failed %v", err)
		return
	}

	var cbbucket *couchbase.Bucket
	cbbucket, err = cbpool.GetBucketWithAuth(*bucketName, mUser, mPassword)

	if err != nil {
		log.Printf("Failed to connect to bucket %v", err)
		return
	}

	log.Printf(" Bucket name %s Bucket %v", *bucketName, cbbucket)
}
