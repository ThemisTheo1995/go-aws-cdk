package main

import (
	"flag"
	"fmt"
	"os"
	"s3share"
)

func main() {
	// Begin CLI
	var from *string
	var bucket *string
	var url string
	var err error

	from = flag.String("file", "", "File to share")
	bucket = flag.String("bucket", "", "Sharing bucket")
	// End CLI

	// Call upload to S3
	s3share.Upload(s3share.Client, from, bucket)
	url, err = s3share.Share(s3share.Client, from, bucket)

	if err != nil {
		fmt.Printf("Error with sharing: %s", err)
		os.Exit(1)
	}

	fmt.Println(url)
}
