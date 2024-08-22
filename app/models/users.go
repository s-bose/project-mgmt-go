package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Password  string    `json:"-" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoUpdateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime:true"`
	Deleted   bool      `json:"deleted" gorm:"not null;default:false"`
}
