package domain

import "time"

type User struct {
	ID         int       `json:"string" gorm:"not null"`
	NIM        string    `json:"NIM" gorm:"not null"`
	Name       string    `json:"name" gorm:"not null"`
	Password   string    `json:"password" gorm:"not null"`
	Department string    `json:"department" gorm:"not null"`
	Role       string    `json:"role" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null"`
}
