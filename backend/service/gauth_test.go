package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YuhriBernardes/gauth-app/service"

	"github.com/YuhriBernardes/gauth-app/database"
	"github.com/YuhriBernardes/gauth-app/model"
)

func TestGauthServiceStartup(t *testing.T) {
	userSample := model.User{Login: "user", Password: "pass"}
	mdb := database.MemoryDatabase{
		Users: map[string]model.User{"someUser": userSample},
	}
	svc := service.GauthService{Db: mdb}

	err := svc.Init()

	assert.Nil(t, err)

}

func TestGauthServiceAuthenticate(t *testing.T) {
	userSample := model.User{Login: "user", Password: "pass"}
	mdb := database.MemoryDatabase{
		Users: map[string]model.User{"someUser": userSample},
	}
	svc := service.GauthService{Db: mdb}

	tests := []struct {
		name               string
		user               model.User
		shouldAuthenticate bool
		authentication     model.Authentication
	}{
		{
			name:               "Authentication with success",
			user:               userSample,
			shouldAuthenticate: true,
			authentication:     model.Authentication{Login: "user", Password: "pass"},
		},
		{
			name:               "Authentication fail -- Wrong login",
			shouldAuthenticate: false,
			authentication:     model.Authentication{Login: "wrongUser", Password: "pass"},
		},
		{
			name:               "Authentication fail -- Wrong password",
			shouldAuthenticate: false,
			authentication:     model.Authentication{Login: "user", Password: "wrongPass"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			usr, err := svc.Authenticate(tt.authentication)

			if tt.shouldAuthenticate {
				assert.Nil(t, err)
				assert.Equal(t, &tt.user, usr)
			} else {
				assert.EqualError(t, model.ErrorUnauthorized, err.Error())
			}
		})

	}

}
