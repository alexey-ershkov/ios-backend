package user

import (
	"context"

	"ios-backend/user/models"
)

type Repository interface {
	Add(c context.Context, newUser models.User) (models.SafeUser, error)
	GetByID(c context.Context, id int) (models.SafeUser, error)
}
