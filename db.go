package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Model
type User struct {
	gorm.Model        //Model comes with ID field, hence not needed in struct
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
}

var DB *gorm.DB

//var err error
const DSN = "root:password@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

//Connect to MySQL Database using gorm
func initialMigration() {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	db.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
