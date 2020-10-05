package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/YuhriBernardes/gauth-app/model"

	"github.com/YuhriBernardes/gauth-app/service"
)

func TestMockServiceAuthentication(t *testing.T) {
	svc := service.MockService{
		Identities: map[string]string{"user1": "password1"},
	}

	tests := []struct {
		Name                string
		Authentication      model.Authentication
		ExpectAuthenticated bool
		AuthenticatedUser   *model.User
	}{
		{
			Name:                "Should be authenticated successfully",
			Authentication:      model.Authentication{Login: "user1", Password: "password1"},
			ExpectAuthenticated: true,
			AuthenticatedUser:   &model.User{Login: "user1", Password: "password1"},
		},
		{
			Name:                "Wrong password authentication",
			Authentication:      model.Authentication{Login: "user1", Password: "12341234"},
			ExpectAuthenticated: false,
		},
		{
			Name:                "Wrong user authentication",
			Authentication:      model.Authentication{Login: "inexistentUser", Password: "password1"},
			ExpectAuthenticated: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			usr, err := svc.Authenticate(tt.Authentication)
			if tt.ExpectAuthenticated {
				assert.Nil(t, err)
				assert.Equal(t, tt.AuthenticatedUser, usr)
			} else {
				assert.EqualError(t, err, model.ErrorUnauthorized.Error())
			}
		})

	}

}
