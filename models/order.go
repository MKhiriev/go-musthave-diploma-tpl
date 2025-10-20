package models

import "time"

type Order struct {
	Order      string    `gorm:"column:order"`
	StatusId   int64     `gorm:"column:status_id"`
	Accrual    float64   `gorm:"column:accrual"`
	UploadedAt time.Time `gorm:"column:uploaded_at"`
}

type Status struct {
	StatusId int64  `gorm:"column:status_id"`
	Name     string `gorm:"column:name"`
}
