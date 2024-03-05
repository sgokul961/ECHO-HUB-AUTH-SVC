package interfacesU

import "github.com/sgokul961/echo-hub-auth-svc/pkg/models"

type AdminUseCase interface {
	AdminSignup(admin models.AdminSignupRequest) (int64, error)
	AdminLogin(admin models.AdminLogin) (models.AdminLoginResposne, error)
	BlockUser(user_id int64) error
}
