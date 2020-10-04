package server

import (
	"github.com/YuhriBernardes/gauth-app/service"
	"github.com/YuhriBernardes/gauth-app/validation"

	"github.com/YuhriBernardes/gauth-app/model"

	"github.com/gin-gonic/gin"
)

var fakeUsers = map[string]bool{
	"edoraoff": true,
	"irar":     false,
}

type Router struct {
	service service.Service
}

type RequestError struct {
	Message string `json:"errMsg"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func (r Router) Authenticate(c *gin.Context) {
	reqBody := model.Authentication{
		Login:    c.PostForm("login"),
		Password: c.PostForm("password"),
	}

	if err := validation.ValidateAuthentication(reqBody); err != nil {
		c.JSON(400, RequestError{Message: err.Error()})
	}

	// if _, err := reqBody.validate(); err != nil {
	// 	c.JSON(http.StatusBadRequest, RequestError{Message: err.Error()})
	// 	return
	// }

	// if v, ok := fakeUsers[reqBody.UserName]; !ok || !v {
	// 	c.Status(http.StatusUnauthorized)
	// 	return
	// }

	// c.JSON(http.StatusOK, AuthenticateResponse{Token: token.GenerateSha512(timestamp, reqBody.UserName)})
}
