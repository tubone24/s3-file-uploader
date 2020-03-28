package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

type S3 struct {
	client *s3.S3
	session *session.Session
}

var s3Session = session.Must(session.NewSession(
	&aws.Config{
		Region: aws.String(Region),
		MaxRetries: aws.Int(3),
	},
))

var s3Instance = &S3{client:s3.New(s3Session)}

func GetS3Instance() *S3 {
	return s3Instance
}

func (s *S3) UploadFile(bucket string, filePath string, key string, contentType string) error {
	var file, err = os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	uploader := s3manager.NewUploader(s3Session)
	params := &s3manager.UploadInput{
		ACL: aws.String("public-read"),
		ContentType: aws.String(contentType),
		Bucket: aws.String(bucket),
		Key: aws.String(key),
		Body: file,
	}

	_, err = uploader.Upload(params)
	if err != nil {
		return err
	}

	return nil
}
