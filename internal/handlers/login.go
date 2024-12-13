package handlers

import (
	"github.com/adamlahbib/go-realtimechat-backend/internal/db"
	"github.com/adamlahbib/go-realtimechat-backend/internal/middleware"
	"github.com/adamlahbib/go-realtimechat-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.IncomingUser

	// parse JSON request into user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// connect to database
	db := db.DbConn()
	var dbUser models.User

	// check if user exists
	if err := db.Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}
	// verify the password
	if dbUser.Password != user.Password {
		c.JSON(400, gin.H{"error": "incorrect password"})
		return
	}

	accessToken, err := middleware.GenerateAccessToken()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := middleware.GenerateRefreshToken()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	models.AuthenticatedUser.Useremail = dbUser.Email

	// return user email and tokens
	c.JSON(200, gin.H{"user_email": dbUser.Email, "access_token": accessToken, "refresh_token": refreshToken})
}
