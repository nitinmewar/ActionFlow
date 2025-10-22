package database

import (
	"database/sql"
	"log"
	"orbit/cmd/env"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	postgresDB *gorm.DB
	sqlDB      *sql.DB
)

func Connection() (*gorm.DB, *sql.DB) {
	var err error

	// singleton pattern
	if postgresDB != nil && sqlDB != nil {
		return postgresDB, sqlDB
	}

	// get DSN from env variables
	dsn := env.DBURL.GetValue()

	// open a postgres connection
	postgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err = postgresDB.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	// db configs
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return postgresDB, sqlDB
}
