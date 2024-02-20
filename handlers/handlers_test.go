package handlers

import (
	"book-rental/data"
	"book-rental/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	var err error

	//rootDir := filepath.Join(".", "..", "..")
	if err != nil {
		panic("failed to get current working directory")
	}

	//dbName := filepath.Join(rootDir, "api22222.db")
	dbName := "api22222.db"
	os.Remove(dbName)

	db.DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	user, books := testDataForDb()
	db.DB.AutoMigrate(&data.User{}, &data.Book{})
	fillDb(db.DB, user, books)

	//	db.DB.Create(&books)

	os.Exit(m.Run())
}

func TestFind(t *testing.T) {
	var books []data.Book

	data.FindAllBooks(db.DB, &books)
	fmt.Println("TESTFIND")
	fmt.Println(books)
	if len(books) != 37 {
		panic("failed to connect database")
	}
}

func TestGetAvailableBooks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/get-available-books", nil)

	record := httptest.NewRecorder()

	GetAvailableBooks(record, req)

	resp := record.Result()

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	test := string(body)
	fmt.Println("TEST")
	fmt.Println(test)
	fmt.Println(len(test))
	var books []data.Book
	json.Unmarshal([]byte(test), &books)
	fmt.Println("Unmarshal")
	fmt.Println(len(books))
	fmt.Println(books)
	if len(books) != 37 {
		t.Errorf("Unexpe ted body return")
	}
}

func TestHelloWorldHandler2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	record := httptest.NewRecorder()

	HelloWord(record, req)

	resp := record.Result()

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	want := `"Hello world"`
	if string(body) != want {
		t.Errorf("Unexpe ted body return")
	}
}

func fillDb(db *gorm.DB, user data.User, books []data.Book) {
	db.Create(&user)
	for _, book := range books {
		result := db.Create(&book)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}

func testDataForDb() (data.User, []data.Book) {

	var books []data.Book

	bookNames := []string{
		"The Alchemist", "To Kill a Mockingbird", "1984", "Brave New World", "The Catcher in the Rye",
		"The Great Gatsby", "Moby-Dick", "Pride and Prejudice", "The Hobbit", "Harry Potter and the Sorcerer's Stone",
		"The Lord of the Rings", "The Shining", "One Hundred Years of Solitude", "The Odyssey", "The Hunger Games",
		"The Girl with the Dragon Tattoo", "Sapiens: A Brief History of Humankind", "The Road", "A Game of Thrones",
		"The Da Vinci Code", "The Fault in Our Stars", "The Hitchhiker's Guide to the Galaxy", "The Martian",
		"The Silence of the Lambs", "The Girl on the Train", "Jurassic Park", "The Chronicles of Narnia",
		"The Kite Runner", "The Grapes of Wrath", "The Color Purple", "Dune", "Fahrenheit 451",
		"The Godfather", "The Outsiders", "The Picture of Dorian Gray", "The Secret Garden", "Wuthering Heights",
	}

	for _, name := range bookNames {
		book := data.Book{
			BookID:          0,
			Name:            name,
			DateToReturn:    nil,
			UserID:          0,
			ShippingAddress: "",
		}
		books = append(books, book)
	}

	// Print the result for verification
	for _, book := range books {
		fmt.Printf("BookID: %d, Name: %s, DateToReturn: %v, UserID: %d, ShippingAddress: %s\n",
			book.BookID, book.Name, book.DateToReturn, book.UserID, book.ShippingAddress)
	}

	user := data.User{
		Email:    "test",
		Password: "test",
		Books:    []data.Book{}, // Initialize Books as an empty slice
	}

	return user, books
}

/*import (
	"book-rental/data"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestHelloWorldHandler2(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	record := httptest.NewRecorder()

	HelloWord(record, req)

	resp := record.Result()

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	want := "Hello world1\n"

	if string(body) != want {
		t.Errorf("Unexpe ted body return")
	}
}

func TestDb(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("existing.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&data.User{}, &data.Book{})
	req := httptest.NewRequest(http.MethodGet, "/get-available-books2", nil)

	w := httptest.NewRecorder()

	GetAvailableBooks(w, req)
}
*/
