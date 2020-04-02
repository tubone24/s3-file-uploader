package logic

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
	"io/ioutil"
	"os"
	"strings"
)


func UploadFileToS3(fileType string, gzippedEncodeData string, fileName string) (result string, err error) {
	log.Info("start upload file to s3: " + fileType + "/" + fileName)
	appConfig := config.GetConfig()
	data, err := base64.StdEncoding.DecodeString(strings.Replace(gzippedEncodeData, "data:text/csv;base64,", "", 1))
	if err != nil {
		return "failed decode Data", err
	}
	tempFile, err := ioutil.TempFile(os.TempDir(), "tempfile-")
	if err != nil {
		return "failed create tmpfile", err
	}

	var gzippedEncodeDataBuffer bytes.Buffer
	_, err = gzippedEncodeDataBuffer.Write(data)
	if err != nil {
		log.Error(err)
		return "failed write gzipEncodeData buffer", err
	}
	EncodeDataBuffer := bytes.Buffer{}
	reader, err := gzip.NewReader(&gzippedEncodeDataBuffer)
	if err != nil {
		log.Error(err)
		return "failed gunzip gzipEncodeData buffer", err
	}
	_, err = EncodeDataBuffer.ReadFrom(reader)
	if err != nil {
		log.Error(err)
		return "failed write EncodeDataBuffer", err
	}

	tempFile.Write(EncodeDataBuffer.Bytes())
	log.Debug("Write file to TMP: " + tempFile.Name())
	err = reader.Close()
	if err != nil {
		log.Error(err)
		return "failed reader close", err
	}
	EncodeDataBuffer.Reset()
	gzippedEncodeDataBuffer.Reset()
	defer tempFile.Close()

	s3 := aws.GetS3Instance()
	key := createPrefix(fileType) + "/" + fileName
	err = s3.UploadFile(appConfig.Aws.BucketName, tempFile.Name(), key, ContentTypeCsv)
	if err != nil {
		log.Error(err)
		return "failed put file to S3", err
	}
	log.Info("upload file: " + key)
	return "success", nil
}

func createPrefix(fileType string) string {
	switch fileType {
	case "test1":
		return "test-test"
	default:
		return fileType
	}
}
