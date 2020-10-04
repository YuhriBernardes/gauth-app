package validation

import (
	"errors"

	"github.com/YuhriBernardes/gauth-app/model"
)

func ValidateAuthentication(authentication model.Authentication) error {
	if !RequiredString(authentication.Login) {
		return errors.New("Field login is required")
	}

	if !RequiredString(authentication.Password) {
		return errors.New("Field password is required")
	}

	return nil
}
