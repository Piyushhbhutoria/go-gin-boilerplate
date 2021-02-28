package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/config"
	"github.com/jackc/pgx"
)

var db *pgx.Conn

func Init() {
	c := config.GetConfig()
	c := config.GetConfig()
	var dbURL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.GetString("db.user"), c.GetString("db.password"), c.GetString("db.host"), c.GetString("db.port"), c.GetString("db.dbname"))

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db = conn
	// defer db.Close(context.Background())
}

func GetDB() *pgx.Conn {
	return db
}
