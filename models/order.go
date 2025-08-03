package models

import (
	"time"
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `gorm:"not null" json:"user_id"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}
