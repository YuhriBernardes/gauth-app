package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerRouter interface {
	Authentication(c *gin.Context)
}

type GinServer struct {
	Port   int
	Host   string
	Router ServerRouter
	Server *gin.Engine
}

func (s *GinServer) Init() {
	s.Server = gin.Default()

	routes := s.Server.Group("/api")
	{
		routes.POST("/tokens", s.Router.Authentication)
	}
}

func (s *GinServer) Start() {
	s.Server.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
