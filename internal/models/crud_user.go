package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type UserCRUD struct {
	ID           pgtype.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Username     string      `json:"username"`
	Email        string      `json:"email" gorm:"unique"`
	PasswordHash string      `json:"-"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	IsActive     bool        `json:"is_active" gorm:"default:true"`
	IsAdmin      bool        `json:"is_admin" gorm:"default:false"`
}

func (u *UserCRUD) TableName() string {
	return "user"
}
