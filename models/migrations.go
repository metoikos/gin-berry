package models

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(conn *gorm.DB) {
	err := conn.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to run migrations. \n", err)
	}
}
