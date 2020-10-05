package service

import (
	"github.com/YuhriBernardes/gauth-app/database"
	"github.com/YuhriBernardes/gauth-app/model"
)

type GauthService struct {
	Db database.Database
}

func (svc GauthService) Init() error {
	err := svc.Db.Init()
	return err
}

func (svc GauthService) Authenticate(authentication model.Authentication) (user *model.User, err error) {
	user, err = svc.Db.GetUserByAuthentication(authentication)

	if err == model.ErrorUserNotFound {
		return user, model.ErrorUnauthorized
	} else if err != nil {
		return user, err
	}

	return user, nil
}
