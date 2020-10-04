package model

import "errors"

var (
	ErrorUnauthorized = errors.New("Wrong login or password")
)
