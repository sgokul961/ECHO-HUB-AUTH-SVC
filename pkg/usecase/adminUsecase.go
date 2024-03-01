package usecase

import (
	"errors"
	"fmt"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/config"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/helper"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
	interfaces "github.com/sgokul961/echo-hub-auth-svc/pkg/repository/interface"
	interfacesU "github.com/sgokul961/echo-hub-auth-svc/pkg/usecase/interface"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/utils"
)

type adminRepo struct {
	AdminRepo interfaces.AdminRepo
	config    config.Config
}

func NewAdminUseCse(Repo interfaces.AdminRepo, userRepo interfaces.UserRepo, config config.Config) interfacesU.AdminUseCase {
	return &adminRepo{
		AdminRepo: Repo,
		config:    config,
	}
}
func (u *adminRepo) AdminSignup(admins models.AdminSignupRequest) (int64, error) {
	valid := helper.IsValidEmail(admins.Email)

	if !valid {
		return 0, errors.New("email is not valid")
	}
	validphoneNum := helper.IsValidPhoneNumber(admins.Phonenum)

	if !validphoneNum {
		return 0, errors.New("not valid phone number")
	}
	userexist, err := u.AdminRepo.IsAdminExist(admins.Email)

	if err != nil {
		return 0, err
	}
	if userexist {
		return 0, errors.New("cant register ")
	}
	hashpassword := utils.HashPassword(admins.Password)
	if hashpassword == "" {
		return 0, errors.New("failed to hash password")
	}
	admins.Password = hashpassword
	admins.Is_Admin = true

	id, err := u.AdminRepo.AdminSignup(admins)
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (a *adminRepo) AdminLogin(admin models.AdminLogin) (models.AdminLoginResposne, error) {

	isadmin, err := a.AdminRepo.IsAdminOrNot(admin.Email)

	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	if !isadmin {
		return models.AdminLoginResposne{}, errors.New("not admin cant login")
	}
	fmt.Println("is admin ", isadmin)

	isExist, err := a.AdminRepo.IsAdminExist(admin.Email)

	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	if !isExist {
		return models.AdminLoginResposne{}, errors.New("user dosent exist")
	}
	password, err := a.AdminRepo.GetPassword(admin.Email)
	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	varifyPassword := utils.CheckPasswordHash(admin.Password, password)
	if !varifyPassword {
		return models.AdminLoginResposne{}, errors.New("password mismatch")
	}
	id, err := a.AdminRepo.FindAdminByMail(admin.Email)
	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	email, err := a.AdminRepo.GetEmail(id)
	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	accessToken, err := utils.GenerateToken(id, a.config.JWTSecretKey, "admin")
	if err != nil {
		fmt.Println("err", err)
		return models.AdminLoginResposne{}, err
	}
	refreshTOken, err := utils.GenerateRefreshToken(id, email, a.config.JWTSecretKey)
	if err != nil {
		return models.AdminLoginResposne{}, err
	}
	return models.AdminLoginResposne{
		AccessToken:  accessToken,
		RefreshToken: refreshTOken,
	}, nil

}
