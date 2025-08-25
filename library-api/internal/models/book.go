package models

import "time"

type Book struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null" json:"title"`
	Author     string    `gorm:"not null" json:"author"`
	Descripion string    `json:"descripion"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
