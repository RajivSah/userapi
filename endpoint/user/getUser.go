package user

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/graniticio/granitic/v2/ws"
)

type DBReader interface {
	All(config *aws.Config) (*dynamodb.ScanOutput, error)
}

type GetUserLogic struct {
	DBHandler DBReader
	DBConfig  AWSConfig
}

func (gl *GetUserLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	users, _ := gl.DBHandler.All(gl.DBConfig.Configuration())
	res.Body = unmarshallList(users)
	res.HTTPStatus = 200
}
