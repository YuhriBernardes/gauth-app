package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   int
	Host   string
	Router Router
	Server *gin.Engine
}

func (s *Server) Init() {
	s.Server = gin.Default()

	routes := s.Server.Group("/api")
	{
		routes.POST("/auth", s.Router.Authenticate)
	}
}

func (s *Server) Start() {
	s.Server.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
