package repository

import (
	"errors"
	"fmt"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
	interfaces "github.com/sgokul961/echo-hub-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

// IsAdminExist implements interfaces.AdminRepo.

func NewAdminRepo(db *gorm.DB) interfaces.AdminRepo {
	return &adminDatabase{
		DB: db,
	}
}
func (u *adminDatabase) AdminLogin(admin models.AdminLogin) (bool, error) {
	var res int

	query := `SELECT COUNT(*) FROM users WHERE email=? and password = ? AND is_admin=true`
	err := u.DB.Raw(query, admin.Email, admin.Password).Scan(&res).Error

	if err != nil {
		return false, err
	}
	return res > 0, err

}

func (u *adminDatabase) AdminSignup(admin models.AdminSignupRequest) (int64, error) {
	var id int64

	query := `INSERT INTO users ( username, email, phonenum, password, profile_picture, bio,gender,is_admin,created_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9) RETURNING id`
	err := u.DB.Raw(query,
		admin.Username,
		admin.Email,
		admin.Phonenum,
		admin.Password,
		admin.ProfilePicture,
		admin.Bio,
		admin.Gender,
		admin.Is_Admin,
		admin.CreatedAt,
	).Scan(&id).Error
	fmt.Println("admin role true or not:", admin.Is_Admin)
	if err != nil {
		return 1, err
	}
	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %v", err)
	}
	return id, nil

}
func (u *adminDatabase) IsAdminExist(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM USERS WHERE email=$1 AND is_admin=true`
	err := u.DB.Raw(query, email).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
func (u *adminDatabase) GetPassword(email string) (string, error) {

	var hashedPassword string
	query := `SELECT password FROM users WHERE email=$1`
	err := u.DB.Raw(query, email).Scan(&hashedPassword).Error
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
func (u *adminDatabase) FindAdminByMail(email string) (int64, error) {
	var id int64
	query := `SELECT id FROM users WHERE email=$1`
	err := u.DB.Raw(query, email).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (u *adminDatabase) GetEmail(id int64) (string, error) {
	var email string

	query := `SELECT email FROM users WHERE id=$1`

	err := u.DB.Raw(query, id).Scan(&email).Error
	if err != nil {
		return "", err
	}
	return email, nil
}
func (u *adminDatabase) IsAdminOrNot(email string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM users WHERE email=? AND is_admin=true`

	err := u.DB.Raw(query, email).Scan(&count).Error

	if err != nil {
		return false, errors.New("error checking admin status")
	}

	isAdmin := count > 0
	return isAdmin, nil
}
