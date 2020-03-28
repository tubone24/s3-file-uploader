package logic

import (
	"encoding/base64"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
	"io/ioutil"
	"os"
)


func UploadFileToS3(fileType string, encodeData string, fileName string) (result string, err error) {
	data, err := base64.StdEncoding.DecodeString(encodeData)
	if err != nil {
		return "failed decode Data", err
	}
	tempFile, err := ioutil.TempFile(os.TempDir(), "tempfile-")
	if err != nil {
		return "failed create tmpfile", err
	}

	tempFile.Write(data)
	defer tempFile.Close()

	s3 := aws.GetS3Instance()
	key := createPrefix(fileType) + "/" + fileName
	err = s3.UploadFile(BucketName, tempFile.Name(), key, ContentTypeCsv)
	if err != nil {
		return "failed put file to S3", err
	}
	return "success", nil
}

func createPrefix(fileType string) string {
	if fileType == "test" {
		return "test-test"
	} else {
		return "other"
	}
}
