package main

import (
	"github.com/YuhriBernardes/gauth-app/database"
	"github.com/YuhriBernardes/gauth-app/model"
	"github.com/YuhriBernardes/gauth-app/server"
	"github.com/YuhriBernardes/gauth-app/service"
)

func main() {

	db := database.MemoryDatabase{
		Users: map[string]model.User{
			"user1": model.User{
				Name:     "Administrator",
				Email:    "admin@admin.com",
				Login:    "admin",
				Password: "admin",
			},
			"user2": model.User{
				Name:     "Lucca Yago Matheus Nunes",
				Email:    "luccayagomatheusnunes-75@callan.com.br",
				Login:    "lucca.nunes",
				Password: "easyPass",
			},
		},
	}

	svc := service.GauthService{
		Db: db,
	}

	router := server.Router{
		Service: svc,
	}

	s := server.GinServer{Router: router, Port: 3001}

	s.Init()
	s.Start()
}
