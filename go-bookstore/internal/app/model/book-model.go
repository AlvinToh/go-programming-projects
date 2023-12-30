package model

import (
	config "github.com/alvintoh/go-programming-projects/go-bookstore/internal/platform/database"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

var CreateBook = func(b *Book) *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

var GetAllBooks = func() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

var GetBookById = func(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

var DeleteBook = func(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
