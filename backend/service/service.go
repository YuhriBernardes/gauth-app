package service

import "github.com/YuhriBernardes/gauth-app/model"

type Service interface {
	Init() error
	Authenticate(model.Authentication) (user *model.User, err error)
}
