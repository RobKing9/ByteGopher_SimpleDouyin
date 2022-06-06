package jwtTool

import (
	"ByteGopher_SimpleDouyin/model"
	"fmt"
	"log"
	"testing"
	"time"
)

var user = &model.User{
	UserId:        1111,
	UserName:      "admin",
	Password:      "admin",
	FollowCount:   100,
	FollowerCount: 100,
}

func TestJwtGenerateToken(t *testing.T) {

	token, err := JwtGenerateToken(user, time.Second*60)
	if err != nil {
		log.Fatal("JGT Err:", err)
	}
	fmt.Println("JGT:", token)
}

func TestJwtParseUser(t *testing.T) {

	token, _ := JwtGenerateToken(user, time.Second*60)
	user, err := JwtParseUser(token)
	if err != nil {
		log.Fatal("JPC Err:", err)
	}

	fmt.Printf("JPC:%#v", user)

}
