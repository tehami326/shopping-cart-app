package models

type OrderItem struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	OrderID  uint `gorm:"not null" json:"order_id"`
	ItemID   uint `gorm:"not null" json:"item_id"`
	Quantity int  `gorm:"not null" json:"quantity"`
}
