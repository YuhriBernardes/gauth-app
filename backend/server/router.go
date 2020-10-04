package server

import (
	"net/http"

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
	Service service.Service
}

type RequestError struct {
	Message string `json:"errMsg"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func (r Router) Authentication(c *gin.Context) {
	reqBody := model.Authentication{
		Login:    c.PostForm("login"),
		Password: c.PostForm("password"),
	}

	if err := validation.ValidateAuthentication(reqBody); err != nil {
		c.JSON(400, RequestError{Message: err.Error()})
		return
	}

	if err := r.Service.Authenticate(reqBody); err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
}
