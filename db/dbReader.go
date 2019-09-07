package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (m DBHandler) All() ([]map[string]*dynamodb.AttributeValue, error) {
	svc := dynamodb.New(session.New(m.Configuration()))
	input := &dynamodb.ScanInput{
		TableName: aws.String(m.ModelName),
	}
	result, err := svc.Scan(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result.Items, nil
}
