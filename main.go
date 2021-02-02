package main

import (
	"fmt"
	"github.com/muhammadzhuhry/bwastartup/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// stop program dan print errornya
		log.Fatal(err.Error())
	}

	fmt.Println("success connect to database")

	var users []user.User
	length := len(users)

	fmt.Println(length)

	// penggunaan find di gorm scr otomatis mencari ke table berdasarkan nama struct
	// ex: struct User table users
	db.Find(&users)

	length = len(users)

	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("=============")
	}
}
