package models

import "time"

type Server struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"not null;index"`
	Name      string `gorm:"not null"`
	Hostname  string
	OS        string
	Platform  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
