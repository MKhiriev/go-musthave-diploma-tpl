package models

import "time"

type Order struct {
	OrderId    int64     `gorm:"column:order_id"`
	Number     string    `gorm:"column:number"`
	StatusId   int64     `gorm:"column:status_id"`
	UserId     int64     `gorm:"column:user_id"`
	Accrual    float64   `gorm:"column:accrual"`
	UploadedAt time.Time `gorm:"column:uploaded_at"`
}

func (u Order) TableName() string {
	return "orders"
}
