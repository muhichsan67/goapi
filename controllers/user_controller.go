package controllers

import (
	"flowcamp-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []models.User

	uc.DB.Find(&users)

	c.JSON(201, gin.H{
		"Data": users,
	})
}
