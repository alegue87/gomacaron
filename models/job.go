package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Uuid    string
	Cronsec string
	Plugin  string
	Service string
	Run     bool `gorm:"default:true"`
}

type JobForm struct {
	Uuid    string `form:"uuid" binding:"required"`
	Cronsec string `form:"cronsec" binding:"required"`
	Plugin  string `form:"plugin"`
	Service string `form:"service"`
	Run     string `form:"run:default:true"`
}
