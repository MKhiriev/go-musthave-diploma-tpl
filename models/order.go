package models

type Order struct {
	Number     string      `gorm:"column:number" json:"number"`
	StatusText string      `gorm:"column:status" json:"status"`
	Accrual    float64     `gorm:"column:accrual" json:"accrual,omitempty"`
	UploadedAt RFC3339Time `gorm:"column:uploaded_at" json:"uploaded_at"`

	OrderId  int64 `gorm:"column:order_id;primarykey" json:"-"`
	StatusId int64 `gorm:"column:status_id" json:"-"`
	UserId   int64 `gorm:"column:user_id" json:"-"`
}

func (u Order) TableName() string {
	return "orders"
}
