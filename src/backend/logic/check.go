package logic

import (
	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
)

func (i *Impl)CheckDynamoLastKey() ([]map[string]string, error){
	var ret []map[string]string
	result, err := i.dynamo.Scan(i.appConfig.Aws.DynamoDBLastKeyTable, 10)
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		for key, value := range item {
			test := aws.AttributeValueToString(key, value)
			ret = append(ret, test)
		}
	}
	return ret, nil
}
