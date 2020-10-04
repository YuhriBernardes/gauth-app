package service

import "github.com/YuhriBernardes/gauth-app/model"

type Service interface {
	Authenticate(model.Authentication) (err error)
}
