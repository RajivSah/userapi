package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBHandler struct {
	ModelName string
}

func (m DBHandler) All(config *aws.Config) (*dynamodb.ScanOutput, error) {
	svc := dynamodb.New(session.New(config))
	input := &dynamodb.ScanInput{
		TableName: aws.String(m.ModelName),
	}
	result, err := svc.Scan(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}
