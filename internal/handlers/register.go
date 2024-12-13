package handlers

import (
	"github.com/adamlahbib/go-realtimechat-backend/internal/db"
	"github.com/adamlahbib/go-realtimechat-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.IncomingUser

	// parse JSON request into user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// connect to database
	db := db.DbConn()
	db.AutoMigrate(&models.User{})

	// check if user already exists
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	// create new user
	newUser := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// return success message
	c.JSON(200, gin.H{"message": "user created"})
}
