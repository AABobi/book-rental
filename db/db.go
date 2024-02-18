package db

import (
	"book-rental/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGDB() {
	db, err := gorm.Open(sqlite.Open("api2.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	/*var books []data.Book = []data.Book{
		{BookID: nil, Name: "The Great Gatsby", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "To Kill a Mockingbird", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "1984", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "Brave New World", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "The Catcher in the Rye", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "Lord of the Flies", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "Moby-Dick", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "Pride and Prejudice", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "The Hobbit", DateToReturn: nil, UserID: nil},
		{BookID: nil, Name: "Harry Potter and the Sorcerer's Stone", DateToReturn: nil, UserID: nil},
	}

	user := data.User{
		Email:    "test",
		Password: "test",
		Books:    []data.Book{}, // Initialize Books as an empty slice
	}*/
	DB = db
	// Migrate the schema
	db.AutoMigrate(&data.User{}, &data.Book{})
	/*	db.Create(&user)
		for _, book := range books {
			result := db.Create(&book)
			if result.Error != nil {
				panic(result.Error)
			}
		}*/
}
