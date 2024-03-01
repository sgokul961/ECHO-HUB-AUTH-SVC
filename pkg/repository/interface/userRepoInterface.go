package interfaces

import (
	"github.com/sgokul961/echo-hub-auth-svc/pkg/domain"
)

type UserRepo interface {
	Create(user domain.User) (int64, error)
	IsUserExist(email string) (bool, error)
	FindUserByMail(email string) (int64, error)
	GetPassword(email string) (string, error)
	GetEmail(id int64) (string, error)
	IsAdminOrNot(email string) (bool, error)
	UpdatePassword(email string, newpassword string, id int64) (bool, error)
	BlockUserWithId(email string) (bool, error)
	CheckIfUserBlocked(id int64) (bool, error)
	IsUserExistWithId(id int64) (bool, error)
}
