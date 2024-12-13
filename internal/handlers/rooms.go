package handlers

import (
	"github.com/adamlahbib/go-realtimechat-backend/internal/db"
	"github.com/adamlahbib/go-realtimechat-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	db := db.DbConn()

	db.AutoMigrate(&models.Room{})

	var room models.Room

	if err := c.BindJSON(&room); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	RoomCreator := models.AuthenticatedUser.Useremail

	var newRoom = models.ChatRoom{
		UserEmail: RoomCreator,
		RoomName:  room.RoomName,
	}

	if err := db.Create(&newRoom).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "room created"})
}

func GetRoomByUserEmail(c *gin.Context) {
	db := db.DbConn()

	userEmail := c.Query("email") // get user email from query string parameter

	if userEmail == "" {
		c.JSON(400, gin.H{"error": "email required"})
		return
	}

	var rooms []models.ChatRoom

	// get all rooms for the user
	if err := db.Where("user_email = ?", userEmail).Find(&rooms).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"rooms": rooms})
}
