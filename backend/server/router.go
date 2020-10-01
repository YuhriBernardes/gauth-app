package server

import (
	"net/http"
	"time"

	"github.com/YuhriBernardes/gauth-app/token"

	"github.com/gin-gonic/gin"
)

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

func (r Router) Authenticate(c *gin.Context) {
	reqBody := &AuthenticateRequest{}
	timestamp := time.Now().Unix()

	if err := c.ShouldBindJSON(reqBody); err != nil {
		c.JSON(http.StatusBadRequest, RequestError{Message: "Invalid json body"})
		return
	}

	c.JSON(200, AuthenticateResponse{Token: token.GenerateSha512(timestamp, reqBody.UserName)})
}
