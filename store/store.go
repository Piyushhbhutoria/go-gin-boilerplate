package store

import (
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.LogMessage("error", "DATABASE_URL environment variable is not set")
		os.Exit(1)
	}

	// Configure GORM logger
	gormLogger := gormLogger.Default.LogMode(gormLogger.Info)

	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db = conn
	logger.LogMessage("info", "postgres db connected via GORM")

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.LogMessage("error", "error getting underlying sql.DB: %v", err)
		os.Exit(1)
	}

	if err := sqlDB.Ping(); err != nil {
		logger.LogMessage("error", "error pinging to db: %v", err)
		logger.LogMessage("debug", "reconnecting")
		Init()
	}

	// Note: Database migrations are handled by golang-migrate
	// Run migrations manually using: make up
	logger.LogMessage("info", "GORM initialized successfully - use 'make up' to run migrations")
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			logger.LogMessage("error", "error getting underlying sql.DB for close: %v", err)
			return
		}
		sqlDB.Close()
	}
}
