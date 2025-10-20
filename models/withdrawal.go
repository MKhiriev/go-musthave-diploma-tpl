package models

import "time"

type Withdrawal struct {
	WithdrawalId int64     `gorm:"column:withdrawal_id" json:"-"`
	UserId       int64     `gorm:"column:login" json:"-"`
	Sum          float64   `gorm:"column:sum" json:"sum"`
	ProcessedAt  time.Time `gorm:"column:processed_at" json:"processed_at"`
}
