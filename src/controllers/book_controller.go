package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/willie/BookstoreAPI/src/models"
	"github.com/willie/BookstoreAPI/src/utils"
)

//Book instance
// var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	//fetch all available book
	newBook := models.GetAllBooks()
	//json encode newbook
	res, _ := json.Marshal(newBook)
	//set the content type
	w.Header().Set("Content-Type", "application/json")
	//send HTTP response header with the provided status code.
	w.WriteHeader(http.StatusOK)
	//write the data to the connection as part of an HTTP reply
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	//get the params variable
	params := mux.Vars(r)
	//get the bookId prop.
	bookId := params["bookId"]
	// log.Fatal(bookId)
	//convert bookId to an int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing data")
	}
	//get book data
	bookDetails, _ := models.GetBookById(Id)
	//json encode newbook
	res, _ := json.Marshal(bookDetails)
	//set the content type
	w.Header().Set("Content-Type", "application/json")
	//send HTTP response header with the provided status code.
	w.WriteHeader(http.StatusOK)
	//write the data to the connection as part of an HTTP reply
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	//create a book struct from r and save in createBook
	utils.ParseBody(r, createBook)
	//create a book object and in db
	book := createBook.CreateBook()
	//json encode book
	res, _ := json.Marshal(book)
	//send HTTP response header with the provided status code.
	w.WriteHeader(http.StatusOK)
	//write the data to the connection as part of an HTTP reply
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//get the params variable
	params := mux.Vars(r)
	//get the bookId prop.
	bookId := params["bookId"]
	//convert bookId to an int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing data")
	}
	//delete a book
	deletedBook := models.DeleteBook(Id)
	//{status: "successfully deleted"}
	//json encode book
	res, _ := json.Marshal(deletedBook)
	//set the content type
	w.Header().Set("Content-Type", "application/json")
	//send HTTP response header with the provided status code.
	w.WriteHeader(http.StatusOK)
	//write the data to the connection as part of an HTTP reply
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//a pointer to models.book
	var updatedBook = &models.Book{}
	//create a book struct from r and save in updatedBook
	utils.ParseBody(r, updatedBook)
	//get params variable provide by user
	params := mux.Vars(r)
	//get the bookId prop. from the params
	bookId := params["bookId"]
	//convert bookId to an int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing string Id")
	}
	//get book and gorm.db data
	bookDetails, db := models.GetBookById(Id)

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	//Save update value in database, if the value doesn't have primary key, will insert it
	db.Save(&bookDetails)
	//json encode book
	res, _ := json.Marshal(bookDetails)
	//set the content type
	w.Header().Set("Content-Type", "application/json")
	//send HTTP response header with the provided status code.
	w.WriteHeader(http.StatusOK)
	//write the data to the connection as part of an HTTP reply
	w.Write(res)

}
