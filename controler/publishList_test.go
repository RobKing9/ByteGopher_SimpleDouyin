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

// PublishList API 控制函数测试

//为测试使用创建 *gin.Engine实例
func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPublishList(t *testing.T) {
	r := setupRouter()
	//将项目中的API注册到测试使用的router
	r.GET("/douyin/publish/list/",middleware.JwtAuthWithUserId(),PublishList)

	// 生成测试用的 token
	token,_ := jwtTool.JwtGenerateToken(&model.User{
		UserId: 1,
	},time.Second)

	//向注册的路有发起请求
	req, _ := http.NewRequest("GET", fmt.Sprintf("/douyin/publish/list/?token=%s&user_id=1",token), nil)
	w := httptest.NewRecorder()

	//模拟http服务处理请求
	r.ServeHTTP(w, req)

	// 状态码返回比较
	assert.Equal(t, http.StatusOK, w.Code)
}
