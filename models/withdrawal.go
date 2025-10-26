package models

import "time"

type Withdrawal struct {
	OrderNum    string    `gorm:"column:order_num" json:"order"`
	Sum         float64   `gorm:"column:sum" json:"sum"`
	ProcessedAt time.Time `gorm:"column:processed_at" json:"processed_at"`

	WithdrawalId int64 `gorm:"column:withdrawal_id" json:"-"`
	UserId       int64 `gorm:"column:user_id" json:"-"`
	OrderId      int64 `gorm:"column:order_id" json:"-"`
}

func (u Withdrawal) TableName() string {
	return "withdrawals"
}
