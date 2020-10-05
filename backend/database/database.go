package database

import "github.com/YuhriBernardes/gauth-app/model"

type Database interface {
	Init() error
	GetUserByAuthentication(model.Authentication) (user *model.User, err error)
}
