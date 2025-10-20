package models

type Balance struct {
	BalanceId int64   `gorm:"column:balance_id" json:"-"`
	UserId    int64   `gorm:"column:user_id" json:"-"`
	Current   float64 `gorm:"column:current" json:"current"`
	Withdrawn float64 `gorm:"column:withdrawn" json:"withdrawn"`
}

func (u Balance) TableName() string {
	return "balance"
}
