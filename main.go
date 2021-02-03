package main

import (
	"github.com/muhammadzhuhry/bwastartup/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Tes simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "artist"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
