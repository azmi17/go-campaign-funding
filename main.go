package main

import (
	"go-campaign-funding/handler"
	"go-campaign-funding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3317)/campaign_startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHanlder(userService)

	router := gin.Default()
	api := router.Group("api/v1", userHandler.RegisterUser)

	// routing path API
	api.POST("/users")

	router.Run(":3000")
}
