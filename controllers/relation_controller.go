package controllers

import (
	"flowcamp-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RelationController struct {
	DB *gorm.DB
}

func NewRelationController(db *gorm.DB) *RelationController {
	return &RelationController{DB: db}
}

func (rc *RelationController) CreateProfile(c *gin.Context) {
	var profile models.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := rc.DB.Create(&profile).Error; err != nil {
		c.JSON(400, gin.H{"Error": "Gagal membuat profile!"})
		return
	}

	c.JSON(201, gin.H{
		"Data": profile,
	})
}
