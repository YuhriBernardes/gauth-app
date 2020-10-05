package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YuhriBernardes/gauth-app/database"

	"github.com/YuhriBernardes/gauth-app/model"
)

func TestGetUserByAuthentication(t *testing.T) {
	userSample := model.User{Name: "Osvaldo Renan Leandro da Rosa", Login: "osvaldo.rosa", Password: "I2rzyrd7Nk", Email: "osvaldorenanleandrodarosa_@accardoso.com.br"}
	mdb := database.MemoryDatabase{Users: map[string]model.User{"userId": userSample}}

	tests := []struct {
		Name           string
		Authentication model.Authentication
		ShouldFindUser bool
		FoundUser      *model.User
	}{
		{
			Name:           "Should the user exists",
			Authentication: model.Authentication{Login: userSample.Login, Password: userSample.Password},
			ShouldFindUser: true,
			FoundUser:      &userSample,
		},
		{
			Name:           "Shouldn't find the user - Wrong Login",
			Authentication: model.Authentication{Login: userSample.Login + "salty", Password: userSample.Password},
			ShouldFindUser: false,
		},
		{
			Name:           "Shouldn't find the user - Wrong Password",
			Authentication: model.Authentication{Login: userSample.Login, Password: userSample.Password + "salty"},
			ShouldFindUser: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			t.Log(tt.Authentication)
			user, err := mdb.GetUserByAuthentication(tt.Authentication)

			if tt.ShouldFindUser {
				assert.Nil(t, err)
				assert.Equal(t, tt.FoundUser, user)

			} else {
				assert.EqualError(t, model.ErrorUserNotFound, err.Error(), "Expected a User Not found error")
			}

		})
	}
}
