package controler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"ByteGopher_SimpleDouyin/middleware"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils/jwtTool"
)

func setEngine() *gin.Engine {
	return gin.Default()
}


func TestAction(t *testing.T) {
	r := setEngine()
	r.POST("/douyin/relation/action/",middleware.JwtAuthWithUserId(),Action)

	// 生成测试用的 token
	token,_ := jwtTool.JwtGenerateToken(&model.User{
		UserId: 1,
	},time.Second)

	url := fmt.Sprintf("/douyin/relation/action/?token=%s&user_id=1&to_user_id=2&action_type=1",token)

	req, _ := http.NewRequest("post",url,nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK,w.Code)
}


func TestFollowList(t *testing.T) {
	r := setEngine()
	r.GET("/douyin/relation/follow/list/",middleware.JwtAuthWithUserId(),FollowList)

	// 生成测试用的 token
	token,_ := jwtTool.JwtGenerateToken(&model.User{
		UserId: 1,
	},time.Second)

	url := fmt.Sprintf("/douyin/relation/follow/list/?token=%s&user_id=1",token)

	req, _ := http.NewRequest("get",url,nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK,w.Code)
}


func TestFollowerList(t *testing.T) {
	r := setEngine()
	r.GET("/douyin/relation/follower/list/",middleware.JwtAuthWithUserId(),FollowerList)

	// 生成测试用的 token
	token,_ := jwtTool.JwtGenerateToken(&model.User{
		UserId: 1,
	},time.Second)

	url := fmt.Sprintf("/douyin/relation/follower/list/?token=%s&user_id=1",token)

	req, _ := http.NewRequest("get",url,nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	assert.Equal(t, http.StatusOK,w.Code)
}