// @title Go Gin Boilerplate API
// @version 1.0
// @description A RESTful API built with Go, Gin, and GORM

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /
// @schemes http https

package main

import (
	"github.com/Piyushhbhutoria/go-gin-boilerplate/logger"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/server"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/store"
)

func main() {
	logger.Init()
	store.Init()
	server.Init()
}
