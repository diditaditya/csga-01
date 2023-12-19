package main

import (
	"log"
	"myGram/config"
	"myGram/handlers"
	"myGram/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Reading from .env file")
	}

	dsn := config.GetDBConfig().GetDBUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})

	r := gin.Default()

	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db))
	r.Use(handlers.Authenticated()).GET("/photos", handlers.GetPhotos(db))
	// r.GET("/comments")
	// r.GET("/social-media")

	r.Run()
}
