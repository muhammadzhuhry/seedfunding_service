package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadzhuhry/bwastartup/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	//dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	// stop program dan print errornya
	//	log.Fatal(err.Error())
	//}
	//
	//fmt.Println("success connect to database")
	//
	//var users []user.User

	// penggunaan find di gorm scr otomatis mencari ke table berdasarkan nama struct
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println("=============")
	//}

	router := gin.Default()
	router.GET("/handler", handler)

	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
