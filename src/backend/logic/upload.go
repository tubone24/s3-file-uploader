package logic

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"strings"
)


func (i *Impl)UploadFileToS3(fileType string, gzippedEncodeData string, fileName string) (result string, err error) {
	log.Info("start upload file to s3: " + fileType + "/" + fileName)
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

	key := createPrefix(fileType) + "/" + fileName
	err = i.uploadFile(tempFile.Name(), key)
	if err != nil {
		log.Error(err)
		return "failed put file to S3", err
	}
	log.Info("upload file: " + key)
	return "success", nil
}

func (i *Impl)uploadFile(filename string, key string) error {
	return i.s3.UploadFile(i.appConfig.Aws.BucketName, filename, key, ContentTypeCsv)
}

func createPrefix(fileType string) string {
	switch fileType {
	case "test1":
		return "test-test"
	default:
		return fileType
	}
}
