package domain

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	NIM        string    `json:"NIM" gorm:"unique;not null;type:varchar(100)"`
	Name       string    `json:"name" gorm:"not null;type:varchar(100)"`
	Password   string    `json:"password" gorm:"not null;type:varchar(100)"`
	Major      string    `json:"major" gorm:"not null;type:varchar(100)"`
	Department string    `json:"department" gorm:"not null;type:varchar(100)"`
	Role       string    `json:"role" gorm:"default:USER"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null"`
}
