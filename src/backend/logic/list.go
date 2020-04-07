package logic

import (
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
	"strconv"
	"strings"
)

func (i *Impl)ListObjects (prefix string) (map[string][]map[string]string, error){
	resp, err := i.listObject(prefix)
	if err != nil {
		log.Error(err)
		return map[string][]map[string]string{"fileList": nil}, err
	}
	returnMap := map[string][]map[string]string{"fileList": convertS3FileObjectInfosToMaps(resp)}
	return returnMap, nil
}

func (i *Impl)listObject(prefix string) ([]aws.S3FileObjectInfo, error){
	return i.s3.ListObject(i.appConfig.Aws.BucketName, prefix)
}

func convertS3FileObjectInfosToMaps(s3ObjectInfos []aws.S3FileObjectInfo) []map[string]string{
	var objectList []map[string]string
	for _, item := range s3ObjectInfos {
		if strings.HasSuffix(item.Name, "/") {
			continue
		}
		objectList = append(objectList, map[string]string{
			"name": item.Name,
			"lastModified": item.LastModified,
			"size": strconv.FormatInt(item.Size, 10) + "B",
		})
	}
	return objectList
}
