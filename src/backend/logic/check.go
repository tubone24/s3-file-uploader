package logic

//import (
//	"github.com/tubone24/s3-file-uploader/src/backend/config"
//	"github.com/tubone24/s3-file-uploader/src/backend/utils/aws"
//)
//
//func checkDynamoLastKey() {
//	appConfig := config.GetConfig()
//	dynamoDB := aws.GetDynamoDBInstance()
//	result, err := dynamoDB.Scan(appConfig.Aws.DynamoDBLastKeyTable)
//	if err != nil {
//		return err
//	}
//	for _, item := range result {
//		key, value := range item {
//			test := aws.AttributeValueToString(key, value),
//		}
//	}
//}
