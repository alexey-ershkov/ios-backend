package usecase

import (
	"context"
	"time"

	"gopkg.in/go-playground/validator.v9"
	"ios-backend/src/configs"
	"ios-backend/src/user"
	"ios-backend/src/user/models"
)

type UserUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo user.Repository, timeOut time.Duration) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, contextTimeout: timeOut}
}

func (u UserUsecase) Add(c context.Context, newUser models.User) (models.SafeUser, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	validation := validator.New()
	if err := validation.Struct(newUser); err != nil {
		return models.SafeUser{}, err
	}

	usr, err := u.userRepo.Add(ctx, newUser)
	if err != nil {
		return models.SafeUser{}, configs.ErrUserAlreadyExist
	}
	return usr, err
}

func (u UserUsecase) GetCurrent(c context.Context, id int) (models.SafeUser, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	usr, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return models.SafeUser{}, configs.ErrUserIsNotRegistered
	}
	return usr, err
}
