package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"strconv"
)

type S3 struct {
	client *s3.S3
	session *session.Session
}

//var s3Session = session.Must(session.NewSession(
//	&aws.Config{
//		Region: aws.String(Region),
//		MaxRetries: aws.Int(3),
//	},
//))

var s3Session = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

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

func (s *S3) ListObject(bucket string, prefix string) ([]S3FileObjectInfo, error){
	var objectList []S3FileObjectInfo
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}
	resp, err := s.client.ListObjects(params)
	if err != nil {
		return objectList, err
	}

	for _, item := range resp.Contents {
		o := S3FileObjectInfo{*item.Key, item.LastModified.String(), *item.Size}
		objectList = append(objectList, o)
	}
	return objectList, nil
}

type S3FileObjectInfo struct {
	Name string
	LastModified string
	Size int64
}

func (s *S3FileObjectInfo )ConvertS3FileObjectInfoToMap() map[string]string{
	return map[string]string{"name": s.Name, "lastModified": s.LastModified, "size": strconv.FormatInt(s.Size, 10)}
}
