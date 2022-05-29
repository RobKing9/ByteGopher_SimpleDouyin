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
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(u model.User) (string, error) {
	expirationtime := time.Now().Add(100 * 24 * time.Hour)
	claims := &Claims{
		UserId: u.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
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
