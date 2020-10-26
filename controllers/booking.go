package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"your.import/path/models"
)

var db *gorm.DB
var err error

type CreateBookingInput struct {
	User    string `json:"user" binding:"required"`
	Members string `json:"members" binding:"required"`
}

type UpdateBookingInput struct {
	User    string `json:"user"`
	Members string `json:"members"`
}

func FindAll(c *gin.Context) {
	var booking []models.Booking
	models.DB.Find(&booking)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": booking})
}

func Find(c *gin.Context) {
	var booking models.Booking
	if err := models.DB.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": booking})
}

func Create(c *gin.Context) {
	// Validate input
	var input CreateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	booking := models.Booking{User: input.User, Members: input.Members}
	models.DB.Create(&booking)

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

func Update(c *gin.Context) {
	// Get model if exist
	var booking models.Booking
	if err := models.DB.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&booking).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

func Delete(c *gin.Context) {
	// Get model if exist
	var booking models.Booking
	if err := models.DB.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&booking)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
