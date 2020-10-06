package server

import (
	"fmt"

	"github.com/gin-contrib/cors"

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

	// Dilable CORS
	s.Server.Use(cors.Default())

	routes := s.Server.Group("/api")
	{
		routes.POST("/tokens", s.Router.Authentication)
	}
}

func (s *GinServer) Start() {
	s.Server.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
