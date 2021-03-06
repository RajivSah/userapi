package user

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/graniticio/granitic/v2/ws"
)

type DBReader interface {
	All() ([]map[string]*dynamodb.AttributeValue, error)
}

type GetUserLogic struct {
	DBHandler DBReader
}

func (gl *GetUserLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	users, _ := gl.DBHandler.All()
	usersRes := []User{}
	err := dynamodbattribute.UnmarshalListOfMaps(users, &usersRes)
	if err != nil {
		panic(fmt.Sprintf("failed to put Unmarshall to Map, %v", err))
	}
	res.Body = usersRes
	res.HTTPStatus = 200
}
