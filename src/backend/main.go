package main

import (
	"github.com/tubone24/s3-file-uploader/src/backend/router"
)

func main() {
	r := router.New()
	r.Logger.Fatal(r.Start("0.0.0.0:8585"))
}
