package models

import (
	"github.com/jinzhu/gorm"
	"github.com/willie/BookstoreAPI/src/config"
)

var db *gorm.DB

type Book struct {
	//Model has model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//initialize the db
func init() {
	config.Connect()
	db = config.GetDB()
	//auto migration for given models, will only add missing fields
	db.AutoMigrate()
}

//createBook method... creates a new book object and into db
func (b *Book) CreateBook() *Book {
	//NewRecord check if value's primary key is blank
	db.NewRecord(b)
	//check if table exist and create table accordingly
	if hastable := db.HasTable(b); !hastable {
		db.CreateTable(&b)
	}

	//inserting the value into database
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var newBooks []Book
	//find a records that match given conditions
	db.Find(&newBooks)
	return newBooks
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

//delete a where from db
func DeleteBook(Id int64) Book {
	var newBook Book
	db.Where("ID=?", Id).Delete(newBook)
	return newBook
}
