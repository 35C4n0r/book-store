package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type PurchaseCRUD struct {
	ID        pgtype.UUID   `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    pgtype.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	ISBN      string        `json:"isbn" gorm:"not null"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Inventory InventoryCRUD `json:"-" gorm:"foreignKey:ISBN;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserCRUD      `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *PurchaseCRUD) TableName() string {
	return "purchases"
}
