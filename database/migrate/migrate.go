package migrate

import (
	"log"
	"sample/database"
	models2 "sample/model"
)

func AutoMigrate() {
	db := database.GetDB()
	err := db.AutoMigrate(&models2.User{}, &models2.Role{})
	if err != nil {
		log.Fatal(err)
	}
}
