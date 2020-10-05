package server

import (
	"net/http"
	"time"

	"github.com/YuhriBernardes/gauth-app/token"

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

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func (r Router) Authentication(c *gin.Context) {
	reqBody := model.Authentication{
		Login:    c.PostForm("login"),
		Password: c.PostForm("password"),
	}

	if err := validation.ValidateAuthentication(reqBody); err != nil {
		c.JSON(400, model.RequestError{Message: err.Error()})
		return
	}

	user, err := r.Service.Authenticate(reqBody)

	if err == model.ErrorUnauthorized {
		c.Status(http.StatusUnauthorized)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.RequestError{Message: err.Error()})
	}

	userToken := token.GenerateSha512(time.Now().Unix(), user.Name, user.Email, user.Login, user.Password)

	c.JSON(http.StatusOK, gin.H{"token": userToken})

}
