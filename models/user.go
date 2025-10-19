package models

type User struct {
	Login        string `gorm:"column:login" json:"login"`
	PasswordHash string `gorm:"column:password_hash" json:"password"`
}

func (u User) TableName() string {
	return "users"
}
