package models

import "time"

type Withdrawal struct {
	WithdrawalId int64     `gorm:"column:withdrawal_id" json:"-"`
	UserId       int64     `gorm:"column:login" json:"-"`
	OrderId      int64     `gorm:"column:order_id" json:"-"`
	Sum          float64   `gorm:"column:sum" json:"sum"`
	ProcessedAt  time.Time `gorm:"column:processed_at" json:"processed_at"`
}

func (u Withdrawal) TableName() string {
	return "withdrawals"
}
