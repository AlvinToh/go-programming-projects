package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func GetDB() *gorm.DB {
	return db
}

func SetDB(database *gorm.DB) {
	db = database
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	db.Find(&Books)
	return Books, nil
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(id int64) error {
	if err := db.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
