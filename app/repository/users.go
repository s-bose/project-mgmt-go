package repository

import (
	"github.com/s-bose/project-mgmt-go/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (u *UserRepository) Create(user models.User) error {
	return u.Db.Create(&user).Error
}

func (u *UserRepository) GetUserByEmail(Email string) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{Email: Email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUser(ID *uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{ID: *ID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
