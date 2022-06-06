package controler

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils/jwtTool"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// Info 用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	//数据库
	db := dao.GetDB()
	//参数
	username := c.PostForm("username")
	psw := c.PostForm("password")
	//数据验证
	//密码必须超过六位
	if len(psw) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "密码必须超过六位！",
		})
		return
	}
	//判断手机号是否存在
	var u *model.User
	db.Where("username=?", username).First(&u)
	if len(u.UserName) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "该用户未注册！"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(psw)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码不正确！"})
		return
	}
	//返回token
	// TODO: releaseToken
	token, err := jwtTool.JwtGenerateToken(u, time.Hour*24*365)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "系统异常！"})
		log.Println("系统异常 err=", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"token": token,
		"msg":   "登录成功!!!",
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	//数据库
	db := dao.GetDB()
	//参数
	username := c.PostForm("username")
	psw := c.PostForm("password")
	//数据验证
	//密码必须超过六位
	if len(psw) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "手机号码必须为11位！",
		})
		return
	}
	//密码加密
	hasedpsw, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "加密错误",
		})
		return
	}
	u := model.User{
		UserName:      username,
		Password:      string(hasedpsw),
		FollowCount:   0,
		FollowerCount: 0,
	}

	//创建此用户
	db.Create(&u)

	c.JSON(200, gin.H{
		"msg": "注册成功!!!",
	})

	log.Println(username, psw)
}
