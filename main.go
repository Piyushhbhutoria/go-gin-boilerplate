package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/config"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/db"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
