package repository

import (
	"github.com/s-bose/project-mgmt-go/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// public decl interface
type UserRepository interface {
	Create(user *models.User) error
	GetUser(ID uuid.UUID) (*models.User, error)
	GetUserByEmail(Email string) (*models.User, error)
	DeleteUser(ID uuid.UUID, Hard bool) error
}

// private impl struct
type userRepository struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *userRepository {
	return &userRepository{Db: db}
}

func (u *userRepository) Create(user *models.User) error {
	return u.Db.Create(user).Error
}

func (u *userRepository) GetUserByEmail(Email string) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{Email: Email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetUser(ID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.Db.Where(&models.User{ID: ID}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) DeleteUser(ID uuid.UUID, Hard bool) error {
	var user models.User
	var query *gorm.DB
	if Hard {
		query = u.Db.Where(&models.User{ID: ID}).Delete(&user)
	} else {
		query = u.Db.Where(&models.User{ID: ID}).Update("deleted", true)
	}

	if err := query.Error; err != nil {
		return err
	} else {
		return nil
	}
}
