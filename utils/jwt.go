package utils

import (
	"ByteGopher_SimpleDouyin/model"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

/*
	utils包 工具包
*/

var identitykey = []byte("a_secrect_crect")

type Claims struct {
	UserId int64
	jwt.StandardClaims
	*model.UserModel
}

func ReleaseToken(u *model.UserModel) (string, error) {
	expirationTime := time.Now().Add(100 * 24 * time.Hour)
	claims := &Claims{
		UserId: u.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ByteGopher",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(identitykey)
	if err != nil {
		log.Println("err = ", err.Error())
	}
	return tokenstring, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return identitykey, nil
	})

	return token, claims, err
}
