package jwtTool

/*
	utils 包 工具包 JWT
	主要用于用户的身份验证，以及操作权限

	调用了第三方库 "github.com/dgrijalva/jwt-go" 使用 JWT
*/

import (
	"errors"
	"fmt"
	"time"

	"ByteGopher_SimpleDouyin/model"
	"github.com/dgrijalva/jwt-go"
)

// AppSecret 作为 token 验证时的密钥
//
// viper.GetString会设置这个值(32byte长度)
var AppSecret = "test"

// AppIss JWT Claims 的签发人
//
// 这个值会被viper.GetString重写
var AppIss = "test"

// 自定义payload结构体,不建议直接使用 dgrijalva/jwtTool-go `jwtTool.StandardClaims`结构体.
type userStdClaims struct {
	jwt.StandardClaims
	*model.User
}

// Valid 实现 `type Claims interface` 的 `Valid() error` 方法
//
// 自定义校验内容
func (c userStdClaims) Valid() (err error) {
	// 验证 token 是否过期
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}

	// 核实发行人是否正确
	if !c.VerifyIssuer(AppIss, true) {
		return errors.New("token's issuer is wrong")
	}

	// 其他自定义核实
	if c.UserId < 0 {
		return errors.New("invalid userId in jwt")
	}
	return
}

// JwtGenerateToken 生成 token
//
// @params m		用户的结构体信息
//
// @params d		token 的有效时间长度
//
// @return string	返回 JWT
//
// 例	JwtGenerateToken(m, time.Hour*24*365)
func JwtGenerateToken(m *model.User, d time.Duration) (string, error) {
	m.Password = ""
	expireTime := time.Now().Add(d)

	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.UserId),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}

	// 生成 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)

	// 返回生成的签名字符串
	return token.SignedString([]byte(AppSecret))
}

// JwtParseUser 解析payload的内容,得到用户信息
func JwtParseUser(tokenString string) (*model.User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}

	claims := &userStdClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims /*KeyFunc*/, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil

	})
	if err != nil {
		return nil, err
	}

	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	return claims.User, nil
}
