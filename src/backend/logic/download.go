package logic

import (
	"github.com/labstack/gommon/log"
)

func (i *Impl)DownloadFileToS3(key string) ([]byte, error) {
	fileBytes, err := i.downloadFile(key)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return fileBytes, nil
}

func (i *Impl)downloadFile(key string) ([]byte, error) {
	return i.s3.DownloadFile(i.appConfig.Aws.BucketName, key)
}
