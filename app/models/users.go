package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoUpdateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime:true"`
}
