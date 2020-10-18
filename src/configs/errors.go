package configs

import "errors"

var (
	//JSON errors
	ErrEmptyJSON        = errors.New("empty jsonData field")
	ErrBadJSON          = errors.New("json parsing error")
	ErrBadRequest       = errors.New("bad request data")
	ErrUserAlreadyExist = errors.New("user already exist")
)
