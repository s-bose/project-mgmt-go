package users

import (
	"github.com/s-bose/project-mgmt-go/app/models"
	"github.com/s-bose/project-mgmt-go/app/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repository repository.UserRepository
}

func New(db *gorm.DB) *UserService {
	return &UserService{Repository: repository.UserRepository{Db: db}}
}

func (s *UserService) InsertUser(Email string, Password string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		return nil, err
	}
	user := models.User{Email: Email, Password: string(hash)}
	err = s.Repository.Create(user)

	return &user, err
}

func (s *UserService) GetUserByEmail(Email string) (*models.User, error) {
	user, err := s.Repository.GetUserByEmail(Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) ValidateUser(user *models.User, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
}
