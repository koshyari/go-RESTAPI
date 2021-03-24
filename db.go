package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Defined ORM model
type User struct {
	gorm.Model        //Model comes with ID field, hence not needed in struct
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
}

//Add your MySQL DSN here
const DSN = "<username>:<password>@tcp(127.0.0.1:3306)/<schema_name>?charset=utf8mb4&parseTime=True&loc=Local"

//Connect to MySQL Database using gorm
func initialMigration() {
	DB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}
