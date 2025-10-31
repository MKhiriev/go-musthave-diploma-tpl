package models

type Withdrawal struct {
	OrderNum    string      `gorm:"column:order_num" json:"order"`
	Sum         float64     `gorm:"column:sum" json:"sum"`
	ProcessedAt RFC3339Time `gorm:"column:processed_at" json:"processed_at"`

	WithdrawalId int64 `gorm:"column:withdrawal_id;primarykey" json:"-"`
	UserID       int64 `gorm:"column:user_id" json:"-"`
	OrderID      int64 `gorm:"column:order_id" json:"-"`
}

func (w Withdrawal) TableName() string {
	return "withdrawals"
}
