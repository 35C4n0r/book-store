package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type ReviewCRUD struct {
	ID        uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    pgtype.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	ISBN      string        `json:"isbn" gorm:"not null"`
	Review    string        `json:"review"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Inventory InventoryCRUD `json:"-" gorm:"foreignKey:ISBN;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserCRUD      `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *ReviewCRUD) TableName() string {
	return "reviews"
}
