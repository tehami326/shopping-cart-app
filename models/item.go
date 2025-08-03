package models

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Status    string    `gorm:"not null" json:"status"`
	Price     float64   `gorm:"default:0" json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
