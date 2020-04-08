package main

import (
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/router"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Error(err)
	}
	r := router.New()
	r.Logger.Fatal(r.Start("0.0.0.0:" + appConfig.Basic.Port))
}
