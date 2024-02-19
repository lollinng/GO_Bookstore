package controllers

import (
	"encoding/json"
	"fmt"

	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Declaration of a variable to hold a new book.
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	// writing response
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) // list of books we found from the db
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// get parameter from the query
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0) // concverting datra to string if not
	if err != nil {
		fmt.Println("error while parsing")
	}

	// accessing book using id
	bookDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)

	// writing response
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	// createBook is a new book struct pointer
	//  &models.Book{} creates a new instance of the Book struct and takes its memory
	// address using the & operator. It initializes the struct with zero values for its fields
	// and passes the address to createbook pointer
	createBook := &models.Book{}
	utils.ParseBody(r, createBook) // synching request and variable , // converts string to json object

	b := createBook.CreateBook()
	res, _ := json.Marshal(b) // converting json object to string

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	updateBook := &models.Book{}   // initializing interface to a variable
	utils.ParseBody(r, updateBook) // Parsing json body and interface

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// updating db values from the request
	bookDetails, db := models.GetBookById(Id) //db value

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
