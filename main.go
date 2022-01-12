package main

import (
	"go-contact/config"
	"go-contact/routes"
)

func main() {

	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()
	// close database when appplication is not running
	defer sqlDB.Close()

	r := routes.SetUpRouter(db)
	r.Run()
}
