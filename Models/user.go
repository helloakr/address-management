package models

type UserTable struct {
	Name     string `gorm:"column:Name;"`
	Password string `gorm:"column:Password;"`
}
