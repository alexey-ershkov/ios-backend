package configs

import "errors"

var (
	//JSON errors
	ErrEmptyJSON           = errors.New("empty jsonData field")
	ErrBadJSON             = errors.New("json parsing error")
	ErrBadRequest          = errors.New("bad request data")
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrUserIsNotRegistered = errors.New("user is not registered") // in get curr user
	ErrUserIdIsNotNumber   = errors.New("user id is not number")  // in get curr user
	NoEnvVarError          = errors.New("There is no var in env")
)
