package entity

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string
	Password string
}
