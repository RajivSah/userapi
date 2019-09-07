package db

import (
	"strconv"
	"userapi/endpoint/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (m DBHandler) Update(input user.User) error {
	svc := dynamodb.New(session.New(m.Configuration()))
	updateInput := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#A": aws.String("Age"),
			"#G": aws.String("Gender"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(input.Name),
			},
			":a": {
				N: aws.String(strconv.Itoa(input.Age)),
			},
			":g": {
				S: aws.String(input.Gender),
			},
			":u": {
				S: aws.String(input.Uid),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Uid": {
				S: aws.String(input.Uid),
			},
		},
		ConditionExpression: aws.String("Uid = :u"),
		ReturnValues:        aws.String("ALL_NEW"),
		TableName:           aws.String(m.ModelName),
		UpdateExpression:    aws.String("SET #N = :n, #A = :a, #G = :g"),
	}

	_, err := svc.UpdateItem(updateInput)
	if err != nil {
		return err
	}
	return nil
}
