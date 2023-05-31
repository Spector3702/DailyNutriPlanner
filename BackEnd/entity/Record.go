package entity

import (
	"time"
)

type Record struct {
	Email string    `gorm:"primaryKey;size:60"`
	Name  string    `gorm:"primaryKey;size:60"`
	Sno   string    `gorm:"size:8"`
	Date  time.Time `gorm:"primaryKey;type:datetime"`

	// foreign keys
	PersonalInfo PersonalInfo `gorm:"foreignKey:Email;references:Email;onDelete:CASCADE;onUpdate:CASCADE"`
	Foodstuff    Foodstuff    `gorm:"foreignKey:Name,Sno;references:Name,Sno;onDelete:CASCADE;onUpdate:CASCADE"`
}

func (Record) TableName() string {
	return "record"
}
