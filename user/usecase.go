package user

import (
	"context"

	"ios-backend/user/models"
)

type Usecase interface {
	Add(c context.Context, newUser models.User) (models.SafeUser, error)
	GetCurrent(c context.Context, id int) (models.SafeUser, error)
}
