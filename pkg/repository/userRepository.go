package repository

import (
	"errors"
	"fmt"

	"github.com/sgokul961/echo-hub-auth-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/models"
	interfaces "github.com/sgokul961/echo-hub-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userDatabase{
		DB: db,
	}
}
func (u *userDatabase) Create(user domain.User) (int64, error) {
	var id int64
	query := `
			INSERT INTO users ( username, email, phonenum, password, profile_picture, bio,gender, is_admin,created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9) RETURNING id
		`
	err := u.DB.Raw(query,

		user.Username,
		user.Email,
		user.Phonenum,
		user.Password,
		user.ProfilePicture,
		user.Bio,
		user.Gender,
		user.Is_Admin,
		user.CreatedAt,
	).Scan(&id).Error
	fmt.Println("role is", user.Is_Admin)

	if err != nil {
		return 0, fmt.Errorf("failed to insert user: %v", err)
	}
	return id, nil

}
func (u *userDatabase) IsUserExist(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM USERS WHERE email=$1 AND is_admin=false`
	err := u.DB.Raw(query, email).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
func (u *userDatabase) FindUserByMail(email string) (int64, error) {
	var id int64
	query := `SELECT id FROM users WHERE email=$1`
	err := u.DB.Raw(query, email).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (u *userDatabase) GetPassword(email string) (string, error) {
	var hashedPassword string
	query := `SELECT password FROM users WHERE email=$1`
	err := u.DB.Raw(query, email).Scan(&hashedPassword).Error
	if err != nil {
		return "", err
	}
	return hashedPassword, nil

}
func (u *userDatabase) GetEmail(id int64) (string, error) {
	var email string

	query := `SELECT email FROM users WHERE id=$1`

	err := u.DB.Raw(query, id).Scan(&email).Error
	if err != nil {
		return "", err
	}
	return email, nil
}
func (u *userDatabase) IsAdminOrNot(email string) (bool, error) {
	var isAdmin bool

	query := `SELECT COUNT(*) FROM users WHERE email=? AND is_admin=true`

	err := u.DB.Raw(query, email).Scan(&isAdmin).Error

	if err != nil {
		return false, errors.New("not valid email id for admin")
	}
	return isAdmin, err

}
func (u *userDatabase) UpdatePassword(email string, newpassword string, id int64) (bool, error) {
	fmt.Println("email and password", email, newpassword)
	query := `UPDATE users SET password = ? WHERE id = ? AND email = ?`

	err := u.DB.Exec(query, newpassword, id, email).Error
	fmt.Println("repo", newpassword, email)

	if err != nil {
		return false, errors.New("password  error")
	}
	return true, nil
}

func (u *userDatabase) BlockUserWithId(email string) (bool, error) {

	query := `UPDATE  users SET is_block=true WHERE email=? `

	err := u.DB.Exec(query, email)

	if err != nil {
		return false, errors.New("cant update this user to blocked ")
	}

	return true, nil

}

func (u *userDatabase) IsUserExistWithId(id int64) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM USERS WHERE id=$1 AND is_block=false`
	err := u.DB.Raw(query, id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
func (u *userDatabase) CheckIfUserBlocked(id int64) (bool, error) {

	var isBlocked bool

	query := `SELECT is_block FROM users  WHERE id = ? `

	err := u.DB.Raw(query, id).Scan(&isBlocked).Error

	if err != nil {
		return false, errors.New(" error in databse ")
	}

	return isBlocked, nil
}
func (u *userDatabase) IsUserExistWIthId(id int64) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM USERS WHERE id=$1 `

	err := u.DB.Raw(query, id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}

func (u *userDatabase) FetchDetails(id int64) (models.UserShortDetail, error) {
	var userDetails models.UserShortDetail

	query := `SELECT id, username, profile_picture FROM users WHERE id = ?`
	err := u.DB.Raw(query, id).Scan(&userDetails).Error
	if err != nil {
		return models.UserShortDetail{}, err
	}

	return userDetails, nil
}
