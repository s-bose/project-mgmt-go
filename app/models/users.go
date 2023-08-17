package models

import (
	"github.com/google/uuid"
	"go.starlark.net/lib/time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
