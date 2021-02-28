package server

import "github.com/Piyushhbhutoria/go-gin-boilerplate/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
