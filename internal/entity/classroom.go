package entity

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	ProfessorID uint   `json:"professorID"`
}
