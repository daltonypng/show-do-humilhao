package entity

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model

	ID          uint `gorm:"primaryKey"`
	Name        string
	ProfessorID uint
	Status      int
}
