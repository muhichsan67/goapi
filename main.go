package main

import (
	"flowcamp-api/config"
	"flowcamp-api/controllers"
	"flowcamp-api/middleware"
	"flowcamp-api/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ENV Gagal terbaca")
	}

	router := gin.Default()

	db := config.ConnectDatabase()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})

	// inisialisasi controllers
	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)
	profileController := controllers.NewRelationController(db)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/users", userController.GetUsers)
			protected.POST("/profiles", profileController.CreateProfile)
		}

	}

	router.Run(":8080")
}
