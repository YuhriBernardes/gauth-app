package service

import (
	"github.com/YuhriBernardes/gauth-app/model"
)

// Identites is a map of login(key) and password(value)
type MockService struct {
	Identities map[string]string
}

func (service MockService) Authenticate(authentication model.Authentication) (user *model.User, err error) {

	password, ok := service.Identities[authentication.Login]

	if !ok {
		return user, model.ErrorUnauthorized
	}

	if password != authentication.Password {
		return user, model.ErrorUnauthorized
	}

	return &model.User{Login: authentication.Login, Password: authentication.Password}, nil
}
