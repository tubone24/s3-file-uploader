package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/labstack/gommon/log"
)

type DynamoDB interface {
	Scan(string) ([]map[string]*dynamodb.AttributeValue, error)
}

type DynamoDBImpl struct {
	client *dynamodb.DynamoDB
}

var dynamoSession = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

var dynamoDBInstance = &DynamoDBImpl{client: dynamodb.New(dynamoSession)}

func GetDynamoDBInstance() *DynamoDBImpl {
	return dynamoDBInstance
}

func (d *DynamoDBImpl) Scan(tableName string, limit int64) ([]map[string]*dynamodb.AttributeValue, error) {
	var params *dynamodb.ScanInput
	if limit == -1 {
		params = &dynamodb.ScanInput{
			TableName: &tableName,
		}
	} else {
		params = &dynamodb.ScanInput{
			TableName: &tableName,
			Limit: &limit,
		}
	}

	res, err := d.client.Scan(params)
	if err != nil {
		return nil, err
	}

	data := res.Items
	lastEvaluatedKey := res.LastEvaluatedKey
	for lastEvaluatedKey != nil {
		params := &dynamodb.ScanInput{
			TableName: &tableName,
			ExclusiveStartKey: lastEvaluatedKey,
		}
		res, err := d.client.Scan(params)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		data = append(data, res.Items...)
		lastEvaluatedKey = res.LastEvaluatedKey
	}


	return data, nil
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
