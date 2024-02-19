package config

/*
config, is designed to manage database connections using the GORM library, which is an ORM (Object Relational Mapping) library for Go
*/

// importing GORM (ORM (Object Relational Mapping) library for GO) and mysql dialect
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// declares a package-level variable db of type *gorm.DB, which will hold the connection to the MySQL database.
// purpose of this file is to return db variable which can be used in
// other files
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/world?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = d // giving db variable mysql db instance
}

// used by other parts of the codebase that need to interact with the database.
func GetDB() *gorm.DB {
	return db
}
