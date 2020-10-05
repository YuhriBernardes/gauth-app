package model

import "errors"

var (
	ErrorUnauthorized = errors.New("Wrong login or password")
	ErrorUserNotFound = errors.New("User not found")
)

type RequestError struct {
	Message string `json:"errMsg"`
}
