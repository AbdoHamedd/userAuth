package user

import (
	"Users/database"
	"Users/models"
)

func createModelUser(user *models.User) {
	database.DB.Create(&user)
}

func emailIsExists(email string) (bool, models.User) {
	user := models.User{}
	database.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false, models.User{}
	}
	return true, user
}
func updateUser(user *models.User) {
	database.DB.Save(&user)
}
func userIsFound(id int) (bool, models.User) {
	user := getUserById(id)
	return user.ID != 0, user
}

func getUserById(id int) models.User {
	user := models.User{}
	database.DB.Where("id = ?", id).First(&user)
	return user
}
