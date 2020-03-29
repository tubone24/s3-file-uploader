package logic

import (
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
	"strings"
)

func ListObjects (prefix string) (map[string][]map[string]string, error){
	appConfig := config.GetConfig()
	var objectList []map[string]string
	s3 := aws.GetS3Instance()
	resp, err := s3.ListObject(appConfig.Aws.BucketName, prefix)
	if err != nil {
		log.Error(err)
		return map[string][]map[string]string{"fileList": objectList}, err
	}
	for _, item := range resp {
		if strings.HasSuffix(item.Name, "/") {
			continue
		}
		objectList = append(objectList, item.ConvertS3FileObjectInfoToMap())
	}
	returnMap := map[string][]map[string]string{"fileList": objectList}
	return returnMap, nil
}
