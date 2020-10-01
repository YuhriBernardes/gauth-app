package main

import (
	"fmt"

	"github.com/YuhriBernardes/gauth-app/server"
)

func main() {
	fmt.Println("Salve meu mundão!!!")
	s := server.Server{Router: server.Router{}, Port: 3001}

	s.Init()
	s.Start()
}
