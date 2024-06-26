package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/pb"
	interfacesU "github.com/sgokul961/echo-hub-auth-svc/pkg/usecase/interface"
)

type UserHandler struct {
	usecase      interfacesU.UserUseCase
	adminusecase interfacesU.AdminUseCase
	pb.UnimplementedAuthServiceServer
	//pb.AuthServiceServer

}

func NewUserHandler(use interfacesU.UserUseCase, aduse interfacesU.AdminUseCase) *UserHandler {
	return &UserHandler{
		usecase:      use,
		adminusecase: aduse,
	}
}

// for userRegister
func (u *UserHandler) Register(ctx context.Context, user *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Convert pb.RegisterRequest to domain.User
	if u.usecase == nil {
		return nil, errors.New("user use case is not initialized")
	}

	domainUser := domain.User{
		Email:          user.Email,
		Password:       user.Password,
		Username:       user.Username,
		Phonenum:       user.Phonenum,
		Bio:            user.Bio,
		Gender:         user.Gender,
		ProfilePicture: user.ProfilePicture,
	}

	// Call the usecase.Register method with the converted domain.User
	response, err := u.usecase.Register(domainUser)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{Id: response}, nil
}

// for User Login
func (u *UserHandler) Login(ctx context.Context, user *pb.LoginRequest) (*pb.LoginResponse, error) {
	userLogin := models.UserLogin{
		Email:    user.Email,
		Password: user.Password,
	}
	fmt.Println("userlogin", userLogin)
	userDetails, err := u.usecase.Login(userLogin)
	if err != nil {
		return nil, err
	}
	fmt.Println("user details ", userDetails)
	return &pb.LoginResponse{
		AccessToken:  userDetails.AccessToken,
		RefreshToken: userDetails.RefreshToken,
	}, nil
}

func (a *UserHandler) AdminSignup(ctx context.Context, admin *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {

	domainAdmin := models.AdminSignupRequest{
		Email:          admin.Email,
		Password:       admin.Password,
		Username:       admin.Username,
		Phonenum:       admin.Phonenum,
		ProfilePicture: admin.ProfilePicture,
		Bio:            admin.Bio,
		Gender:         admin.Gender,
	}
	response, err := a.adminusecase.AdminSignup(domainAdmin)
	if err != nil {
		return nil, err
	}
	return &pb.AdminSignupResponse{Id: response}, nil

}
func (a *UserHandler) AdminLogin(ctx context.Context, admin *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	adminLogin := models.AdminLogin{
		Email:    admin.Email,
		Password: admin.Password,
	}
	adminDetails, err := a.adminusecase.AdminLogin(adminLogin)
	if err != nil {
		return nil, err
	}
	return &pb.AdminLoginResponse{
		AccessToken:  adminDetails.AccessToken,
		RefreshToken: adminDetails.RefreshToken,
	}, nil
}
func (a *UserHandler) ResetPassword(ctx context.Context, r *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	fmt.Println("handler pass and email", r.Email, r.Password)
	_, err := a.usecase.UpdatePassword(r.Email, r.Password, r.Id)

	if err != nil {
		return nil, err
	}
	fmt.Println("pb res", r.Email, r.Password)
	return &pb.ResetPasswordResponse{
		Status: http.StatusOK,
	}, nil
}
func (a *UserHandler) CheckUserBlocked(ctx context.Context, r *pb.CheckUserBlockedRequest) (*pb.CheckUserBlockedResponse, error) {
	is_block, err := a.usecase.CheckUserBlocked(r.Id)

	if err != nil {
		return nil, err
	}
	return &pb.CheckUserBlockedResponse{
		Status: is_block,
		Userid: r.Id,
	}, nil
}
func (a *UserHandler) BlockUser(ctx context.Context, r *pb.BlockUserRequest) (*pb.BlockUserResponse, error) {

	err := a.adminusecase.BlockUser(r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.BlockUserResponse{
		Status: http.StatusOK,
		Userid: r.Id,
	}, nil

}
func (a *UserHandler) UnblockUser(ctx context.Context, r *pb.UnblockUserRequest) (*pb.UnblockUserResponse, error) {

	err := a.adminusecase.UnblockUser(r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UnblockUserResponse{
		Status: http.StatusOK,
		Userid: r.Id,
	}, nil

}
func (a *UserHandler) FetchShortDetails(ctx context.Context, r *pb.FetchShortDetailsRequest) (*pb.FetchShortDetailsResponse, error) {
	user, err := a.usecase.FetchDetails(r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.FetchShortDetailsResponse{
		Id:    user.Id,
		Name:  user.Username,
		Image: user.ProfilePicture,
	}, nil

}
