package data

import (
	"book-rental/utils"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	UserID   uint   `json:"user_id" binding:"required" gorm:"primary_key"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required" gorm:"not null"`
	Books    []Book `json:"books"`
}

func CreateNewUser(db *gorm.DB, user *User) {
	var userExist []User
	db.Where("email = ?", user.Email).Find(&userExist)

	if len(userExist) != 0 {
		fmt.Println("tutaj obsługa")
		return
	}
	db.Create(&user)
}

func (u *User) CheckCredentials(db *gorm.DB) bool {
	var user User
	db.Where("email = ?", u.Email).Find(&user)

	return utils.CheckPasswordHash(u.Password, user.Password)
}

func (u *User) FindUser(db *gorm.DB) uint {
	var user User
	db.Where("email = ?", u.Email).Find(&user)

	return user.UserID
}

// test method
func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}
