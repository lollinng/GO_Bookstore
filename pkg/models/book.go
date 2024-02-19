package models

/*
This Go model file defines a Book struct and provides functions to interact with a database table storing book records.

gorm:"": This part of the struct tag is used by the GORM library.
It specifies the options for the field when interacting with the database. In this case, it appears to be empty,
which means GORM will use default settings for this field.

json:"name": This part of the struct tag is used for JSON serialization and deserialization. It specifies how the field
should be represented in JSON format. Here, it indicates that when the struct is marshaled to JSON, the field should be named "name".

init(): This function is automatically called when the package is initialized. It connects to

When AutoMigrate is called:

If the books table (assuming GORM's default table name convention) does not exist in the database,
GORM will create it based on the structure of the Book model.

If the books table already exists in the database, GORM will inspect the table structure and make any necessary changes
to ensure it matches the fields of the Book model. This might involve adding new columns, modifying column types, or altering constraints.

*/

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

// package-level variable db of type *gorm.DB, which is used to interact with the database.
var db *gorm.DB

// gorm.Model,adds fields like ID, CreatedAt, UpdatedAt, and DeletedAt to manage records in the database.
type Book struct {
	gorm.Model
	Name        string ` gorm:"" json:"name" `
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// to intialize the db ,  automatically called when the package is initialized. It connects to
func init() {

	// connecting using config package we wrote
	config.Connect()
	db = config.GetDB()

	// This function is used to automatically create or update database tables to match the structure of the given model.
	db.AutoMigrate(&Book{})
}

// models to access the database
// gorm writes our query and provides an abstract layer for us
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) //gorm NewRecord function
	db.Create(&b)   // creating new record
	return b        // return same object
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return &book
}
