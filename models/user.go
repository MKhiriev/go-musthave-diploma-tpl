package models

type User struct {
	UserId   int64  `gorm:"column:user_id" json:"-"`
	Login    string `gorm:"column:login" json:"login"`
	Password string `gorm:"column:password_hash" json:"password"`
}

func (u User) TableName() string {
	return "users"
}
