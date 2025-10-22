package models

import "time"

type Order struct {
	OrderId    int64     `gorm:"column:order_id"`
	Number     string    `gorm:"column:number"`
	StatusId   int64     `gorm:"column:status_id"`
	Accrual    float64   `gorm:"column:accrual"`
	UploadedAt time.Time `gorm:"column:uploaded_at"`
}

type Status struct {
	StatusId int64  `gorm:"column:status_id"`
	Name     string `gorm:"column:name"`
}

func (u Order) TableName() string {
	return "orders"
}

func (u Status) TableName() string {
	return "statuses"
}
