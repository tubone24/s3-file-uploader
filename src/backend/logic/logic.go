package logic

import (
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
)

type Logic interface {
	DeleteFileToS3(fileKey string) error
	deleteObject(fileKey string) error
	ListObjects(prefix string) (map[string][]map[string]string, error)
	listObject(prefix string) ([]aws.S3FileObjectInfo, error)
	UploadFileToS3(fileType string, gzippedEncodeData string, fileName string) (result string, err error)
	uploadFile(filename string, key string) error
}

type Impl struct {
	appConfig *config.Config
	s3 *aws.S3Impl
}

func New() *Impl {
	impl := Impl{
		appConfig: config.GetConfig(),
		s3: aws.GetS3Instance(),
	}
	return &impl
}
