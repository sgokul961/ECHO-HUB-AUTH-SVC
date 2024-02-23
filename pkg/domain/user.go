package domain

import (
	"time"
)

type User struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Username       string    `json:"username" validate:"required,min=8,max=24"`
	Email          string    `json:"email" validate:"email,required"`
	Phonenum       string    `json:"phonenum" validate:"required,len=10"`
	Password       string    `json:"password" validate:"required,min=8,max=16"`
	ProfilePicture string    `json:"profile_picture"`
	Gender         string    `json:"gender"`
	Bio            string    `json:"bio"`
	Is_Admin       bool      `json:"is_admin" gorm:"default:false"`
	IsBlock        bool      `json:"is_block" gorm:"default:false"`
	CreatedAt      time.Time `json:"created_at"`
}
type Image struct {
	ID     int64  `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId int64  `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserId"`
	Image  string `json:"image"`
}

//gorm:"check:role IN ('admin', 'user')"
