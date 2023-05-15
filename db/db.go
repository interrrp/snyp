// Package db contains a global database connection and
// all schemas.
package db

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/interrrp/snyp/util"
)

const (
	// DefaultDSN is the default DSN to connect to when the
	// POSTGRES_DSN environment variable doesn't exist.
	DefaultDSN = "host=localhost user=postgres password=postgres database=snyp sslmode=disable"

	// RetryDelay is the time to wait between
	// connection retries.
	RetryDelay = 3 * time.Second
)

var (
	// DB is the global database connection.
	DB *gorm.DB

	// DSN is the DSN/connection string.
	DSN = util.EnvOr("POSTGRES_DSN", DefaultDSN)

	// Config is the GORM config.
	Config = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
)

// init initializes the database connection.
func init() {
	// Keep trying to connect to the database.
	err := errors.New("")
	for err != nil {
		DB, err = gorm.Open(postgres.Open(DSN), Config)
		if err != nil {
			log.Println("Could not connect to database:", err.Error())
		}

		time.Sleep(RetryDelay)
	}
}

// init registers all schemas.
func init() {
	DB.AutoMigrate(&Snippet{})
}
