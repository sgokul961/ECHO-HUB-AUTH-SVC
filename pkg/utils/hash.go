package utils

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return "cant hash"
	}
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
func VerifyPassword(requestPassword, dbPassword string) bool {
	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	fmt.Println("req", requestPassword)
	return requestPassword == dbPassword
}
