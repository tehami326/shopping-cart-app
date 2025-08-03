package models

import (
	"time"
)

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items"`
}
