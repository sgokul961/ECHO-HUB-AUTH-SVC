package models

import (
	"mime/multipart"
	"time"
)

// response when user tries to login
type RegisterRequestBody struct {
	Email          string          `json:"email"`
	Password       string          `json:"password"`
	Username       string          `json:"username"`
	Phonenum       string          `json:"phonenum"`
	ProfilePicture *multipart.Form `json:"profile_picture"`
	Bio            string          `json:"bio"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLoginResponse struct {
	RefreshToken string
	AccessToken  string
}
type AdminLogin struct {
	Email    string
	Password string
}
type AdminLoginResposne struct {
	RefreshToken string
	AccessToken  string
}
type AdminSignupRequest struct {
	Username       string    `json:"username" validate:"required,min=8,max=24"`
	Email          string    `json:"email" validate:"email,required"`
	Phonenum       string    `json:"phonenum" validate:"required,len=10"`
	Password       string    `json:"password" validate:"required,min=8,max=16"`
	ProfilePicture string    `json:"profile_picture"`
	Gender         string    `json:"gender"`
	Bio            string    `json:"bio"`
	Is_Admin       bool      `json:"is_admin" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
}
