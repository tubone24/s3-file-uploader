package logic

import (
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
)

func DownloadFileToS3(key string) ([]byte, error) {
	appConfig := config.GetConfig()
	s3 := aws.GetS3Instance()
	fileBytes, err := s3.DownloadFile(appConfig.Aws.BucketName, key)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return fileBytes, nil
}

