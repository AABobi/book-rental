package data

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookID       uint       `json:"book_id" gorm:"primaryKey;autoIncrement;not null"`
	Name         string     `json:"name" gorm:"not null;not null"`
	DateToReturn *time.Time `json:"date_to_return"`
	UserID       *uint      `json:"user_id" gorm:"foreignKey:UserID"`
}

func FindBooks(db *gorm.DB, condition string) []Book {
	var books []Book
	db.Where(condition).Find(&books)
	return books
}

func FindAllBooks(db *gorm.DB, books *[]Book) {
	db.Find(&books)
	//return books
}

func (b *Book) RentBook(db *gorm.DB, userId *uint) {
	db.Where("book_id = ?", b.BookID).First(&b)
	currentTime := time.Now()

	// Add 2 weeks to the current time
	twoWeeksLater := currentTime.Add(2 * 7 * 24 * time.Hour)

	//b.DateToReturn = &twoWeeksLater

	if err := db.Model(&b).Update("date_to_return", twoWeeksLater).Error; err != nil {
		panic("Failed to update user")
	}

	b.UserID = userId
	fmt.Printf("user id in book %v and thje vaue %v", b.UserID, *userId)
	//b.BookID = 1
	if err := db.Model(&b).Update("user_id", *userId).Error; err != nil {
		panic("Failed to update user")
	}
	fmt.Println("User updated successfully:", b)

}
