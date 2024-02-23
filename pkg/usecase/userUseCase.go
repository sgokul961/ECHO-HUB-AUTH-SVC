package usecase

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/config"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
	interfaces "github.com/sgokul961/echo-hub-auth-svc/pkg/repository/interface"
	interfacesU "github.com/sgokul961/echo-hub-auth-svc/pkg/usecase/interface"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/utils"
)

type userUseCase struct {
	repo interfaces.UserRepo

	config config.Config
}

func NewUserUseCase(repository interfaces.UserRepo, config config.Config) interfacesU.UserUseCase {
	return &userUseCase{
		repo:   repository,
		config: config,
	}
}

func (u *userUseCase) Register(user domain.User) (int64, error) {

	//validate email
	valid, err := u.ValidateEMail(user.Email)
	if err != nil {
		return 0, err
	}
	if !valid {
		return 0, errors.New("email is in error format")
	}
	validNum := u.IsValidPhoneNumber(user.Phonenum)

	if err != nil {
		return 0, err
	}
	if !validNum {
		return 1, errors.New("phone number is not valid")
	}
	userexist, err := u.repo.IsUserExist(user.Email)

	if err != nil {
		return 0, err
	}
	if userexist {
		return 0, errors.New("user already exist")
	}
	hashpassword := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashpassword

	//user.Role = "user"

	id, err := u.repo.Create(user)

	if err != nil {
		return 0, err
	}
	return id, nil

}
func (u *userUseCase) Login(user models.UserLogin) (models.UserLoginResponse, error) {

	//checking if user exist with the same email id

	userExist, err := u.repo.IsUserExist(user.Email)
	fmt.Println("user exist", userExist)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	if !userExist {
		return models.UserLoginResponse{}, errors.New("user not exist")

	}

	//getting dbpassword from user databse with the email id

	password, err := u.repo.GetPassword(user.Email)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	fmt.Println("user passw", password)

	//varifiying the hashed password from db with the  entered password

	verifypassword := utils.CheckPasswordHash(user.Password, password)
	if !verifypassword {
		return models.UserLoginResponse{}, errors.New("password mismatch")
	}
	fmt.Println("var pass", verifypassword)

	//fetching  id with the email associated with it
	id, err := u.repo.FindUserByMail(user.Email)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	fmt.Println("findmail", id)

	email, err := u.repo.GetEmail(id)
	if err != nil {
		return models.UserLoginResponse{}, err

	}

	fmt.Println("getmail", email)

	//creating access token

	accessToken, err := utils.GenerateToken(id, u.config.JWTSecretKey, "user")
	if err != nil {
		fmt.Println("err", err)
		return models.UserLoginResponse{}, err
	}
	fmt.Println("acctoekn", accessToken)
	// return models.UserLoginResponse{AccessToken: accessToken}, nil

	//creating refresh token
	refreshTOken, err := utils.GenerateRefreshToken(id, email, u.config.JWTSecretKey)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	fmt.Println("refreshtok", refreshTOken)
	return models.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTOken,
	}, nil

}
func (u *userUseCase) ValidateEMail(email string) (bool, error) {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, email)
	if err != nil {
		return false, err
	}
	return match, nil

}
func (u *userUseCase) IsValidPhoneNumber(phoneNumber string) bool {
	regex := `^[0-9]{10}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match
}
