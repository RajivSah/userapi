package user

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/graniticio/granitic/v2/ws"
)

type DBWriter interface {
	Write(config *aws.Config, input interface{}) error
}

type AddUserLogic struct {
	DBHandler DBWriter
	DBConfig  AWSConfig
}

func (al *AddUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, u *User) {
	err := al.DBHandler.Write(al.DBConfig.Configuration(), *u)
	if err != nil {
		res.HTTPStatus = 422
		return
	}
	res.HTTPStatus = 201
}
