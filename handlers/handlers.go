package handlers

import (
	"book-rental/data"
	"book-rental/db"
	"book-rental/helpers"
	"book-rental/utils"
	"fmt"

	"net/http"
)

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

func GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	condition := "user_id IS NULL"
	books := data.FindBooks(db.DB, condition)
	_ = helpers.WriteJSON(w, http.StatusOK, books)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := helpers.ReadJSON(w, r, &user)
	if err != nil {
		//errorJSON(w, err)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Println("Hashed password error")
	}

	user.Password = hashedPassword
	data.CreateNewUser(db.DB, &user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := helpers.ReadJSON(w, r, &user)

	if err != nil {
		fmt.Println("err")
		return
	}

	passwordCorrect := user.CheckCredentials(db.DB)

	if !passwordCorrect {
		fmt.Println("Wrong password")
		return
	}

	userId := user.FindUser(db.DB)

	token, err := utils.GenerateToken(user.Email, &userId)

	if err != nil {
		fmt.Println("Generate token error")
		return
	}

	response := AuthResponse{
		Email: user.Email,
		Token: token,
	}
	helpers.WriteJSON(w, 200, response)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := data.GetAllUsers(db.DB)
	helpers.WriteJSON(w, 200, users)
}

func RentBook(w http.ResponseWriter, r *http.Request) {
	var book data.Book
	helpers.ReadJSON(w, r, &book)
	userId, ok := r.Context().Value("myKey").(*uint)

	if !ok {
		http.Error(w, "Failed to retrieve user ID from context", http.StatusInternalServerError)
		return
	}

	book.RentBook(db.DB, userId)
	//fmt.Println(*book.BookID, book.Name)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []data.Book
	data.FindAllBooks(db.DB, &books)
	helpers.WriteJSON(w, 200, books)
}
