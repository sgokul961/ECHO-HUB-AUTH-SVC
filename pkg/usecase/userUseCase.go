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
	"golang.org/x/crypto/bcrypt"
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

	// if err != nil {
	// 	return 0, err
	// }
	if !validNum {
		return 0, errors.New("phone number is not valid")
	}
	userexist, err := u.repo.IsUserExist(user.Email)

	if err != nil {
		return 0, err
	}
	if userexist {
		return 0, errors.New("user already exist")
	}
	hashpassword := utils.HashPassword(user.Password)
	if hashpassword == "" {
		return 0, errors.New("failed to hash password")
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
	fmt.Println("enterd password", user.Password)

	//varifiying the hashed password from db with the  entered password

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return models.UserLoginResponse{}, errors.New("password mismatch")
	}
	fmt.Println("password is ", user.Password, password)

	// if err != nil {
	// 	return models.UserLoginResponse{}, err
	// }

	//fmt.Println("var pass", verifypassword)

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
func (u *userUseCase) UpdatePassword(email string, newpassword string, id int64) (bool, error) {
	//checking if the user exist in the incoming email
	user, err := u.repo.IsUserExist(email)
	if err != nil {
		return false, err
	}
	if !user {
		return false, errors.New("user doesn't exist")

	}
	//getting the email id with the incoming email
	userid, err := u.repo.FindUserByMail(email)
	fmt.Println("user id ", id)
	if err != nil {
		return false, err
	}

	//if the user id is not same as the id inthe context it willbe returned .
	if userid != id {

		return false, errors.New("user ID doesn't match the email provided")

	}
	//hash the new password
	hash := utils.HashPassword(newpassword)
	fmt.Println("usecase emailand password", email, newpassword)
	passwordUpdate, err := u.repo.UpdatePassword(email, hash, id)

	if err != nil {
		return false, err
	}
	return passwordUpdate, nil

}
func (u *userUseCase) CheckUserBlocked(id int64) (bool, error) {
	exsist, err := u.repo.IsUserExistWIthId(id)

	if err != nil {
		return false, err
	}
	if !exsist {
		return false, errors.New("user id invalid")
	}

	isBlocked, err := u.repo.CheckIfUserBlocked(id)

	// Print status based on isBlocked
	fmt.Println("isblock", isBlocked)
	if err != nil {
		return false, err // Pass the actual error returned by the repository
	}

	return isBlocked, nil
}
