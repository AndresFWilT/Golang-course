package model

import "errors"

var (
	ErrPersonCannotBeNil = errors.New("person cannot be nil")
	ErrPersonIdNotFound  = errors.New("person not found")
)
