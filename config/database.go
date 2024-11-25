package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	logger := GetLogger("initdb")
	// create and connect to database
	gorm.Open(postgres.Open(), &gorm.Config{})
}
