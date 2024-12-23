package models

import "Users/database"

func Migrate() {
	database.DB.AutoMigrate(&User{})
}
