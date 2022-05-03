package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const dsn = "postgresql://gin_service:test123@localhost:5432/gin_service"

func Init() *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database.", err)
	}

	DB = db

	return DB
}
