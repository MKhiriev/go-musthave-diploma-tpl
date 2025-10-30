package models

type Balance struct {
	Current   float64 `gorm:"column:current" json:"current"`
	Withdrawn float64 `gorm:"column:withdrawn" json:"withdrawn,omitempty"`

	BalanceId int64 `gorm:"column:balance_id;primarykey" json:"-"`
	UserId    int64 `gorm:"column:user_id" json:"-"`
}

func (u Balance) TableName() string {
	return "balance"
}
