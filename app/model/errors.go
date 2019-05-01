package model

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrUserameConflict     = errors.New("Username already exist")
	ErrEmailConflict       = errors.New("Email already exist")
	ErrPasswordConflict    = errors.New("Password is wrong")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)
