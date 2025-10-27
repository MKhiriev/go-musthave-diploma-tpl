package models

const (
	NEW        = "NEW"
	PROCESSING = "PROCESSING"
	INVALID    = "INVALID"
	PROCESSED  = "PROCESSED"
)

type Status struct {
	StatusId int64  `gorm:"column:status_id"`
	Name     string `gorm:"column:name"`
}

func (u Status) TableName() string {
	return "statuses"
}
