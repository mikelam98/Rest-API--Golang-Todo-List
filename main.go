package main

import (
	"Golang-Crud/module/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/todo_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
	}
	log.Println("Connecting to database successfully")

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/items", transport.HandleCreateNewItem(db))
		v1.GET("/items", transport.HandleListItem(db))
		v1.PUT("/items/:id", transport.HandleUpdateItem(db))
		v1.GET("/items/:id", transport.HandleFindItem(db))
		v1.DELETE("/items/:id", transport.HandleDeleteItem(db))
	}

	if err := router.Run(); err != nil {
		log.Println(err)
	}
}
