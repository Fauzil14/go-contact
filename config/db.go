package config

import (
	"fmt"
	"go-contact/models"
	"go-contact/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// configuration for connecting to database

func ConnectDatabase() *gorm.DB {

	username := utils.GetEnv("DB_USER", "fauzil")
	password := utils.GetEnv("DB_PASS", "password")
	database := utils.GetEnv("DB_NAME", "db_contact")
	host := utils.GetEnv("DB_HOST", "tcp(127.0.0.1:3306)") // tcp(localhost:port)

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&models.Contact{})

	return db
}
