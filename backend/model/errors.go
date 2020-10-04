package model

import "errors"

var (
	ErrorUnauthorized = errors.New("Wrong login or password")
)

type RequestError struct {
	Message string `json:"errMsg"`
}
