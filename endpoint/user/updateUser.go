package user

import (
	"context"

	"github.com/graniticio/granitic/v2/ws"
)

type DBUpdater interface {
	Update(input User) error
}

type UpdateUserLogic struct {
	DBHandler DBUpdater
}

func (al *UpdateUserLogic) ProcessPayload(ctx context.Context, req *ws.Request, res *ws.Response, u *User) {
	err := al.DBHandler.Update(*u)
	if err != nil {
		res.HTTPStatus = 422
		return
	}
	res.HTTPStatus = 200
}
