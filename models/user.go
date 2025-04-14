package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Token    string
}

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Token    string `form:"token"`
}
