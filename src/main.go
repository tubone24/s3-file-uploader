package main

import (
	"github.com/tubone24/s3-file-uploader/src/router"
)

func main() {
	r := router.New()
	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
