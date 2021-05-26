package models

type ClassName struct {
	Name string `gorm:"column:ClassName;"`
	Time string `gorm:"column:CreatTime;"`
}
