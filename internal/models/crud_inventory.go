package models

import "time"

type InventoryCRUD struct {
	ISBN        string    `json:"isbn" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Content     []byte    `json:"content" gorm:"type:bytea"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *InventoryCRUD) TableName() string {
	return "inventory"
}
