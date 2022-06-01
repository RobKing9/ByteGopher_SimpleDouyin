package controler

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils/jwtTool"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Info 用户信息
func Info(c *gin.Context) {

}

// Login 用户登录
func Login(ctx *gin.Context) {
	userReq := &LoginRequest{
		UserName: ctx.Query("username"),
		Password: ctx.Query("password"),
	}

	ctx.JSON(http.StatusOK,LoginService(userReq))
}

// Register 用户注册
func Register(c *gin.Context) {

}



func LoginService(req *LoginRequest) UserLoginResponse {
	// 查询用户是否存在
	usr,err := dao.NewUserModel().SearchUser(req.UserName,req.Password)
	if err != nil {
		return UserLoginResponse{
			Response:model.Response{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		}
	}

	//fmt.Printf("%#v",usr)

	// 生成token
	token, err := jwtTool.JwtGenerateToken(&usr, time.Second*120)
	if err != nil {
		return UserLoginResponse{
			Response:model.Response{
				StatusCode: model.SCodeFalse,
				StatusMsg: "login failed",
			},
			UserId: usr.UserId,
		}
	}

	// 将 token 与 user 的关系存入缓存中
	// 明确缓存过期时间
	// 用户主动退出等登录后删除
	//cache.Cache.CacheToken(token,usr)
	dao.GetTestData(usr.UserId)

	return UserLoginResponse{
		Response:model.Response{
			StatusCode: model.SCodeSuccess,
		},
		UserId: usr.UserId,
		Token: token,
	}
}


type UserLoginResponse struct {
	model.Response
	UserId		int64	`json:"user_id,omitempty"`
	Token		string	`json:"token"`		// 用户鉴权 token
}


type LoginRequest struct {
	UserName	string	`json:"user_name"`
	Password	string	`json:"password"`
}
