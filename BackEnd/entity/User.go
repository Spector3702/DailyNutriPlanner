package entity

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key:auto_increment" json:"id"`
	Nickname  string    `gorm:"type:varchar(255);not null" json:"nickname"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"size:varchar(255);not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
