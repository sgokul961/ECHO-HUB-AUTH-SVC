package interfacesU

import (
	"github.com/sgokul961/echo-hub-auth-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
)

type UserUseCase interface {
	Register(user domain.User) (int64, error)
	// ValidateUser(user models.User) error
	Login(models.UserLogin) (models.UserLoginResponse, error)
	ValidateEMail(email string) (bool, error)
	IsValidPhoneNumber(phoneNumber string) bool
}
