package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type CartCRUD struct {
	ID        uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    pgtype.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	ISBN      string        `json:"isbn" gorm:"not null"`
	Quantity  int           `json:"quantity" gorm:"not null;default:1"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      UserCRUD      `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Inventory InventoryCRUD `json:"-" gorm:"foreignKey:ISBN;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *CartCRUD) TableName() string {
	return "cart"
}
