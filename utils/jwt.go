package utils

import (
	"ByteGopher_SimpleDouyin/model"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
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

//// JwtParseUser 解析payload的内容,得到用户信息
//func JwtParseUser(tokenString string) (*model.User, error) {
//	if tokenString == "" {
//		return nil, errors.New("no token is found in Authorization Bearer")
//	}
//
//	claims := &Claims{}
//
//	_, err := jwt.ParseWithClaims(tokenString, claims /*KeyFunc*/, func(token *jwt.Token) (interface{}, error) {
//
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(AppSecret), nil
//
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	err = claims.Valid()
//	if err != nil {
//		return nil, err
//	}
//
//	return claims.User, nil
//}
//
//// 自定义校验内容
//func (c userStdClaims) Valid() (err error) {
//	// 验证 token 是否过期
//	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
//		return errors.New("token is expired")
//	}
//
//	// 核实发行人是否正确
//	if !c.VerifyIssuer(AppIss, true) {
//		return errors.New("token's issuer is wrong")
//	}
//
//	// 其他自定义核实
//	if c.UserId < 0 {
//		return errors.New("invalid userId in jwt")
//	}
//	return
//}
