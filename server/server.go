package server

import (
	"log"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	err := r.Run(config.GetString("server.port"))
	if err != nil {
		log.Println("error running server => ", err)
	}
}
