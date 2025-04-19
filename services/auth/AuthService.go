package auth

import (
	"github.com/shannon3335/story-server/internal/types"
	"gorm.io/gorm"
)

type AuthService interface {
	SignupUser(types.User) error
	Login(Username string, Password string) (bool, error)
	GetUser(Email string) (*types.User, error)
}

type authService struct {
	DB *gorm.DB
}

func NewAuthService(DB *gorm.DB) AuthService {
	return &authService{
		DB: DB,
	}
}

func (a *authService) SignupUser(user types.User) error {
	hashedPass, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	result := a.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *authService) Login(Username string, Password string) (bool, error) {
	var hashedPass string
	err := a.DB.Select("password").Where("username = ?", Username).First(&hashedPass).Error
	if err != nil {
		return false, err
	}
	return CompareWithString(Password, []byte(hashedPass)), nil
}

func (a *authService) GetUser(Email string) (*types.User, error) {
	var user types.User
	err := a.DB.Where("email = ?", Email).Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}
