package models

const (
	NEW        = "NEW"
	PROCESSING = "PROCESSING"
	INVALID    = "INVALID"
	PROCESSED  = "PROCESSED"
)

type Status struct {
	StatusID int64  `gorm:"column:status_id;primarykey"`
	Name     string `gorm:"column:name"`
}

func (u Status) TableName() string {
	return "statuses"
}
