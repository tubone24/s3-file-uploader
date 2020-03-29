package logic

import (
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
)

func DeleteFileToS3(fileKey string) error {
	appConfig := config.GetConfig()
	s3 := aws.GetS3Instance()
	err := s3.DeleteObject(appConfig.Aws.BucketName, fileKey)
	if err != nil {
		return err
	}
	return nil
}
