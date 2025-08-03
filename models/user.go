package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Token     string    `gorm:"default:''" json:"token"`
	Cart      Cart      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
}
