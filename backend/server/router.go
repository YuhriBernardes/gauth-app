package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/YuhriBernardes/gauth-app/validation"

	"github.com/YuhriBernardes/gauth-app/token"

	"github.com/gin-gonic/gin"
)

var fakeUsers = map[string]bool{
	"edoraoff": true,
	"irar":     false,
}

type Router struct {
}

type RequestError struct {
	Message string `json:"errMsg"`
}

type AuthenticateRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func (req *AuthenticateRequest) validate() (ok bool, err error) {
	if !validation.RequiredString(req.UserName) {
		return false, errors.New("Field userName is required")
	}

	if !validation.RequiredString(req.Password) {
		return false, errors.New("Field password is required")
	}

	return true, nil

}

func (r Router) Authenticate(c *gin.Context) {
	reqBody := &AuthenticateRequest{}
	timestamp := time.Now().Unix()

	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, RequestError{Message: "Invalid json body"})
		return
	}

	if _, err := reqBody.validate(); err != nil {
		c.JSON(http.StatusBadRequest, RequestError{Message: err.Error()})
		return
	}

	if v, ok := fakeUsers[reqBody.UserName]; !ok || !v {
		c.Status(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, AuthenticateResponse{Token: token.GenerateSha512(timestamp, reqBody.UserName)})
}
