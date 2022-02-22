package models

import (

	// go-sql-driver/mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SetupDB : initializing mysql database
func ConfigDB() *gorm.DB {
	// USER := "root"
	// PASS := "password"
	// HOST := "localhost"
	// PORT := "3306"
	// DBNAME := "bookCRUD"
	// URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	// db, err := gorm.Open("mysql", URL)
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=testdbgo sslmode=disable password=asdf1234")
	if err != nil {
		panic(err.Error())
	}
	return db
}
