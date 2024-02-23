package interfaces

import (
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
)

type AdminRepo interface {
	AdminLogin(admin models.AdminLogin) (bool, error)
	AdminSignup(admin models.AdminSignupRequest) (int64, error)
	IsAdminExist(email string) (bool, error)
	GetPassword(email string) (string, error)
	FindAdminByMail(email string) (int64, error)
	GetEmail(id int64) (string, error)
	IsAdminOrNot(email string) (bool, error)
}
