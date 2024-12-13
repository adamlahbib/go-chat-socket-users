package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

type IncomingUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChatRoom struct {
	gorm.Model
	UserEmail string `json:"useremail"`
	RoomName  string `json:"roomname"`
}

type Room struct {
	RoomName string `json:"roomname"`
}

type Message struct {
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type authenticatedUser struct {
	Useremail string `json:"email"`
}

var AuthenticatedUser authenticatedUser
