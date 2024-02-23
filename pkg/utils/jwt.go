package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtClaims struct {
	jwt.StandardClaims
	Id   int64
	Role string
}

func GenerateToken(id int64, key string, role string) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:   id,
		Role: role,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 400).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("key", key)
	signedToken, err = token.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userid int64, email string, key string) (string, error) {

	refreshTokenExpiresAt := time.Now().Add(time.Hour * 24 * 7).Unix() // Refresh token expires in 7 days
	refreshTokenClaims := jwt.MapClaims{
		"exp":  refreshTokenExpiresAt,
		"sub":  userid,
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(key))
	return refreshTokenString, err
}
