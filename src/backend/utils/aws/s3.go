package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

type S3 interface {
	UploadFile(bucket string, filePath string, key string, contentType string) error
	ListObject(bucket string, prefix string) ([]S3FileObjectInfo, error)
	DownloadFile(bucket string, key string) ([]byte, error)
	DeleteObject(bucket string, key string) error
}

type S3Impl struct {
	client *s3.S3
	session *session.Session
}

type S3FileObjectInfo struct {
	Name string
	LastModified string
	Size int64
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

var s3Instance = &S3Impl{client:s3.New(s3Session)}

func GetS3Instance() *S3Impl {
	return s3Instance
}

func (s *S3Impl) UploadFile(bucket string, filePath string, key string, contentType string) error {
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

func (s *S3Impl) ListObject(bucket string, prefix string) ([]S3FileObjectInfo, error){
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
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		formattedTime := item.LastModified.In(jst).Format(time.RFC3339)
		o := S3FileObjectInfo{*item.Key, formattedTime, *item.Size}
		objectList = append(objectList, o)
	}
	return objectList, nil
}

func (s *S3Impl) DownloadFile(bucket string, key string) ([]byte, error) {
	downloader := s3manager.NewDownloader(s3Session)
	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(key),
	}

	fileBuf := &aws.WriteAtBuffer{}

	_, err := downloader.Download(fileBuf, params)
	if err != nil {
		return nil, err
	}

	fileBytes := fileBuf.Bytes()

	return fileBytes, nil
}

func (s *S3Impl) DeleteObject(bucket string, key string) error {
	params := &s3.DeleteObjectInput{
		Bucket:aws.String(bucket),
		Key:aws.String(key),
	}

	_, err := s.client.DeleteObject(params)
	if err != nil {
		return err
	}

	return nil
}
