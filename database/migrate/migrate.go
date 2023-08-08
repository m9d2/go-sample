package migrate

import (
	"log"
	"sample/app/model"
	"sample/database"
)

func AutoMigrate() {
	db := database.GetDB()
	err := db.AutoMigrate(&model.User{}, &model.Role{})
	if err != nil {
		log.Fatal(err)
	}
}
