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
	// Dilable CORS
	s.Server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	routes := s.Server.Group("/api")
	{
		routes.POST("/auth", s.Router.Authenticate)
	}
}

func (s *Server) Start() {
	s.Server.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
}
