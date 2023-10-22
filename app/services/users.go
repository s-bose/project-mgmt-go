package users

import (
	"time"

	"github.com/s-bose/project-mgmt-go/app/models"
	"github.com/s-bose/project-mgmt-go/app/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	InsertUser(Email string, Password string) (*models.User, error)
	GetUserByEmail(Email string) (*models.User, error)
	ValidateUser(User *models.User, Password string) error
}

type userService struct {
	Repository repository.UserRepository
}

func New(db *gorm.DB) *userService {
	return &userService{Repository: repository.New(db)}
}

func (s *userService) InsertUser(Name string, Email string, Password string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 10)
	if err != nil {
		return nil, err
	}
	user := models.User{Name: Name, Email: Email, Password: string(hash), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err = s.Repository.Create(&user)
	return &user, err
}

func (s *userService) GetUserByEmail(Email string) (*models.User, error) {
	user, err := s.Repository.GetUserByEmail(Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ValidateUser(user *models.User, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
}
