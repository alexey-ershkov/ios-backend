package user

import (
	"context"

	"ios-backend/src/user/models"
)

type Usecase interface {
	Add(c context.Context, newUser models.User) (models.SafeUser, error)
	GetCurrent(c context.Context, id int) (models.SafeUser, error)
	GetUserByEmailAndPassword(c context.Context, email string, password string) (models.SafeUser, error)
}
