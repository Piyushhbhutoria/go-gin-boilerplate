package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/config"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/logger"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.GetString("db.host"), c.GetInt("db.port"), c.GetString("db.user"), c.GetString("db.password"), c.GetString("db.dbname"))

	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db = conn
	// defer db.Close()
}

func GetDB() *sql.DB {
	if err := db.Ping(); err != nil {
		logger.LogMessage("error", "error pinging to db", err)
		logger.LogMessage("debug", "reconnecting", nil)
		Init()
	}
	return db
}
