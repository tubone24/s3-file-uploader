package logic

import "github.com/labstack/gommon/log"

func (i *Impl)DeleteFileToS3(fileKey string) error {
	err := i.deleteObject(fileKey)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (i *Impl)deleteObject(fileKey string) error {
	return i.s3.DeleteObject(i.appConfig.Aws.BucketName, fileKey)
}
