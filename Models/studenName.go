package models

type StudenTables struct {
	ClassName string `gorm:"column:ClassName;"`
	Name      string `gorm:"column:Name;"`
	StudenID  string `gorm:"column:StudenID;"`
	QQNumble  string `gorm:"column:QQNumble;"`
	Email     string `gorm:"column:Email;"`
	Address   string `gorm:"column:Address;"`
	Password  string `gorm:"column:Password;"`
}
