package entity

import (
	"time"
)

type Record struct {
	Email string `gorm:"primaryKey;size:60"`
	Name  string `gorm:"primaryKey;size:60"`
	Sno   string `gorm:"primaryKey;size:8"`
	Date  time.Time

	// foreign keys
	PersonalInfo PersonalInfo `gorm:"foreignKey:Email;references:Email;onDelete:CASCADE;onUpdate:CASCADE"`
	Foodstuff    Foodstuff    `gorm:"foreignKey:Name,Sno;references:Name,Sno;onDelete:CASCADE;onUpdate:CASCADE"`
}
