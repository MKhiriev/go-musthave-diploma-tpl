package models

const (
	NEW        = "NEW"
	Processing = "PROCESSING"
	Invalid    = "INVALID"
	Processed  = "PROCESSED"
)

type Status struct {
	StatusId int64  `gorm:"column:status_id"`
	Name     string `gorm:"column:name"`
}

func (u Status) TableName() string {
	return "statuses"
}
