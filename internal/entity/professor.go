package entity

type Professor struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string
	Password string
}
