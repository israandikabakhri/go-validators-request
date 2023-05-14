package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.core/go-validator-request/config/app"
)

func init() {
	app.LoadEnvVariabels()
}

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("MARIADB")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return database
}
