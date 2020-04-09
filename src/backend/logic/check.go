package logic

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/imdario/mergo"
)

func (i *Impl)CheckDynamoLastKey() ([]map[string]string, error){
	var ret []map[string]string
	result, err := i.dynamo.Scan(i.appConfig.Aws.DynamoDBLastKeyTable, 10)
	if err != nil {
		return nil, err
	}
	c := 0
	var tm map[string]string
	for _, item := range result {
		for key, value := range item {
			kv := AttributeValueToString(key, value)
			err = mergo.Merge(&tm, kv)
			if err != nil {
				return nil, err
			}
			if c == len(i.appConfig.DB.KeyList) -1 {
				c = 0
				tmn := tm
				ret = append(ret, tmn)
				tm = map[string]string{}
			} else {
				c += 1
			}
		}
	}
	return ret, nil
}

func AttributeValueToString(key string, value *dynamodb.AttributeValue) map[string]string {
	stringItem := map[string]string {}
	stringItem[key] = value.String()
	if value.N != nil {
		stringItem[key] = *value.N
	} else if value.S != nil {
		stringItem[key] = *value.S
	}
	return stringItem
}
