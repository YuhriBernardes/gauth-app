package main

import (
	"fmt"

	"github.com/YuhriBernardes/gauth-app/server"
	"github.com/YuhriBernardes/gauth-app/service"
)

func main() {
	fmt.Println("Salve meu mundão!!!")

	router := server.Router{
		Service: service.MockService{
			Identities: map[string]string{
				"admin": "admin",
				"user":  "1234",
			},
		},
	}

	s := server.GinServer{Router: router, Port: 3001}

	s.Init()
	s.Start()
}
