package user

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type DBWriter interface {
	Write(input interface{}) error
}

type AddUserLogic struct {
	DBHandler DBWriter
}

func (al *AddUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, u *User) {
	err := al.DBHandler.Write(*u)
	if err != nil {
		res.HTTPStatus = 422
		return
	}
	res.HTTPStatus = 201
}
