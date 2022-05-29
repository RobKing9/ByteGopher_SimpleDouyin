package controler

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"strconv"

	//"encoding/json"
	"github.com/gin-gonic/gin"
)

//func UnmarshalRespModel(data []byte) (RespModel, error) {
//	var r RespModel
//	err := json.Unmarshal(data, &r)
//	return r, err
//}

//func (r *RespModel) Marshal() ([]byte, error) {
//	return json.Marshal(r)
//}

// upload video的返回
type RespModel struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 视频列表的返回
type RespVideoList struct {
	StatusCode int64         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	VideoList  []model.Video `json:"video_list"` // 用户发布的视频列表
}

// 根据视频信息，生成视频的结构体对象
func makeVideo(id int64, author model.User, playUrl string, coverURl string, favoriteCount int64, commentCount int64, isFavorite bool, title string) model.Video {
	return model.Video{
		Id:            id,
		Author:        author,
		PlayUrl:       playUrl,
		CoverUrl:      coverURl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         title,
	}
}

// 上传用户视频
func VideoUpload(c *gin.Context) {
	data, _ := c.FormFile("data")
	token := c.PostForm("token")
	title := c.PostForm("title")

	id := utils.RandRangeIn(10000000, 99999999) // 随机数生成视频ID
	var FavoriteCount int64 = 0
	var CommentCount int64 = 0
	IsFavorite := false

	var testVideo model.Video
	var user model.User

	db := dao.GetDB()

	// TODO: 判断上传的视频文件是否为mp4格式

	// TODO: 对token做鉴权 根据token获取用户
	user = GetUserFromToken(token)

	// 判断id是否存在 若已存在则重新生成
	for {
		db.Where("id = ?", id).First(&testVideo)
		if testVideo.Id == 0 {
			break
		}
		id = utils.RandRangeIn(10000000, 99999999)
	}
	// 视频先保存到本地
	c.SaveUploadedFile(data, "/tmp/${id}.mp4")

	//TODO: 获取封面url 目前封面暂定一个固定的图
	coverUrl := "http://cdn1.pic.y1ng.vip/uPic/IMG_3567.JPG"

	// 上传至七牛云
	retKey, err := utils.QiniuUpload("video/${id}.mp4", "/tmp/${id}.mp4")

	// 无论远程oss上传是否成功 都删除本地文件
	utils.DeleteFile("/tmp/${id}.mp4")

	// 如果上传失败则返回
	if err != nil {
		res := RespModel{
			StatusCode: -1,
			StatusMsg:  err.Error(),
		}
		c.JSON(200, res)
		return
	}

	video := makeVideo(int64(id), user, retKey, coverUrl, FavoriteCount, CommentCount, IsFavorite, title)

	// 插入数据到数据库
	result := db.Create(&video)

	// 判断是否插入成功
	if result.Error != nil {
		res := RespModel{
			StatusCode: -1,
			StatusMsg:  result.Error.Error(),
		}
		c.JSON(200, res)
		return
	}

	// 返回成功
	res := RespModel{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	c.JSON(200, res)
}

// 获取用户发布的视频列表
func VideoList(c *gin.Context) {
	token := c.PostForm("token")
	userID := c.PostForm("user_id")
	userIDint, _ := strconv.ParseInt(userID, 10, 64)

	var videoList []model.Video
	db := dao.GetDB()

	// TODO: 对token做鉴权 根据token获取用户
	var user model.User
	user = GetUserFromToken(token)

	// 根据token获取的用户的user id和传入的user id 不一致，直接return 防止越权漏洞
	if user.UserId != userIDint {
		res := RespVideoList{
			StatusCode: -1,
			StatusMsg:  "user id not match",
			VideoList:  videoList,
		}
		c.JSON(200, res)
		return
	}

	result := db.Where("user_id = ?", userID).Find(&videoList)

	// 查询失败
	if result.Error != nil {
		res := RespVideoList{
			StatusCode: -1,
			StatusMsg:  result.Error.Error(),
			VideoList:  videoList,
		}
		c.JSON(200, res)
		return
	}

	// 查询成功
	res := RespVideoList{
		StatusCode: 0,
		StatusMsg:  "success",
		VideoList:  videoList,
	}
	c.JSON(200, res)
}
