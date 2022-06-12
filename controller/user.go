package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserController interface {
	Info(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type userController struct {
	userDao dao.UserDao
}

func NewUserController() UserController {
	return &userController{
		userDao: dao.NewUserDao(),
	}
}

func (controller userController) Info(c *gin.Context) {
	// 认证失败
	flag, _ := c.Get("flag")
	if !flag.(bool) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": -1,
			"status_msg":  "请先登录!!!",
			"user":        nil,
		})
		log.Println("请先登录！")
		return
	}
	// 获取用户id
	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	// 通过用户id获取用户
	user, err := controller.userDao.GetCommonUserByID(user_id)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": -1,
			"status_msg":  "用户id不存在!!!",
			"user":        nil,
		})
		log.Println("用户id不存在！")
		return
	}

	// TODO: 当前用户是否follow了这个user， is_follow字段
	// TODO: 等关注功能写完 再回来写这里

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "successful",
		"user":        user,
	})
}

// Login 用户登录
func (controller userController) Login(c *gin.Context) {
	//参数
	username := c.Query("username")
	pwd := c.Query("password")

	//数据验证
	//密码必须超过六位
	// if len(psd) < 6 {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": "密码必须超过六位！",
	// 	})
	// 	log.Println("密码必须超过六位！")
	// 	return
	// }
	//判断手机号是否存在

	u, err := controller.userDao.GetUserByName(username)
	// db.Where("user_name=?", username).First(&u)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": -1,
			"status_msg":  "该用户未注册!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("该用户未注册")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "密码不正确!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("密码不正确!")
		return
	}
	//返回token
	token, err := utils.ReleaseToken(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  "系统异常!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("系统异常 err=", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "登录成功!!!",
		"user_id":     u.UserID,
		"token":       token,
	})
	log.Println("登录成功！")
}

// Register 用户注册
func (controller userController) Register(c *gin.Context) {
	//参数
	username := c.Query("username")
	pwd := c.Query("password")

	// 判断用户名是否存在
	uu, err := controller.userDao.GetUserByName(username)
	if err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  "系统异常!!!",
				"user_id":     0,
				"token":       0,
			})
			log.Println("系统异常")
			return
		}
	}
	fmt.Printf("%+v\n", uu)
	if uu != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  "该用户名已存在!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("用户名已存在")
		return
	}

	//数据验证
	//密码必须超过六位
	if len(pwd) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": -1,
			"status_msg":  "密码必须超过六位!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("密码必须超过六位！")
		return
	}

	//密码加密
	hasedpsw, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  "加密错误!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("加密错误")
		return
	}
	rand.Seed(time.Now().UnixNano())

	id := rand.Int63() // 生成比较大的随机数
	u := model.UserModel{
		UserID:        id,
		UserName:      username,
		Password:      string(hasedpsw),
		FollowCount:   0,
		FollowerCount: 0,
		//IsFollow:      false,
	}
	// db.AutoMigrate(model.User{})
	//创建此用户
	// db.Create(&u)
	controller.userDao.AddUserModel(&u)
	//返回token
	token, err := utils.ReleaseToken(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  "系统异常!!!",
			"user_id":     0,
			"token":       0,
		})
		log.Println("系统异常 err=", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "注册成功!!!",
		"user_id":     id,
		"token":       token,
	})

	log.Printf("注册成功!用户id为：%d, 用户名为：%s,密码是：%s token:%s", id, username, pwd, token)
}
