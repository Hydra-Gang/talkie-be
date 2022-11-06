package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Hydra-Gang/talkie-be/configs"
	"github.com/Hydra-Gang/talkie-be/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	db, err := configs.SetupDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	// migrates all models
	db.AutoMigrate(&models.User{})

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
