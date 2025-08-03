package models

type CartItem struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	CartID   uint `gorm:"not null" json:"cart_id"`
	ItemID   uint `gorm:"not null" json:"item_id"`
	Quantity int  `gorm:"not null" json:"quantity"`
}
