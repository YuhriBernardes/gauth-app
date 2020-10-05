package database

import "github.com/YuhriBernardes/gauth-app/model"

// They key in Users is the User ID
type MemoryDatabase struct {
	Users map[string]model.User
}

func (mdb *MemoryDatabase) Init() error {
	return nil
}

func (mdb *MemoryDatabase) GetUserByAuthentication(authentication model.Authentication) (user *model.User, err error) {

	for _, u := range mdb.Users {
		if u.Login == authentication.Login && u.Password == authentication.Password {
			return &u, nil
		}
	}

	return user, model.ErrorUserNotFound
}
