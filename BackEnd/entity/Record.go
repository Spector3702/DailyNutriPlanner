package entity

import (
	"time"
)

type Record struct {
	Email string    `gorm:"primaryKey;size:60"`
	Name  string    `gorm:"primaryKey;size:60"`
	Sno   string    `gorm:"size:8"`
	Date  time.Time `gorm:"primaryKey;type:datetime"`
}

func (Record) TableName() string {
	return "record"
}
