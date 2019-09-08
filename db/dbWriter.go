package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

func (m DBHandler) Write(input interface{}) error {
	svc := dynamodb.New(session.New(m.Configuration()))
	uid, _ := uuid.NewUUID()
	marshalledInput, err := dynamodbattribute.MarshalMap(input)
	*marshalledInput["Uid"] = dynamodb.AttributeValue{S: aws.String(uid.String())}
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
		return err
	}
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(m.ModelName),
		Item:      marshalledInput,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
		return err
	}
	return nil
}
