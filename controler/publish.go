package controler

import (
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"github.com/gin-gonic/gin"
)
import "encoding/json"

func UnmarshalRespModel(data []byte) (RespModel, error) {
	var r RespModel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RespModel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RespModel struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func makeVideoModel(id int64, author User, playUrl string, coverURl string, favoriteCount int64, commentCount int64, isFavorite bool, title string) model.VideoModel {
	return model.VideoModel{
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

func VideoUpload(c *gin.Context) {
	data, _ := c.FormFile("data")
	token := c.PostForm("token")
	title := c.PostForm("title")
	id := utils.RandRangeIn(10000000, 99999999) // 随机数生成视频ID
	var FavoriteCount int64 = 0
	var CommentCount int64 = 0
	IsFavorite := false

	// TODO: 判断上传的视频文件是否为mp4格式

	// TODO: 对token做鉴权 根据token获取用户
	var user User
	user = GetUserFromToken(token)

	//TODO: 判断id是否存在 若已存在则重新生成

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

	video := makeVideoModel(int64(id), user, retKey, coverUrl, FavoriteCount, CommentCount, IsFavorite, title)

	// TODO: 插入数据到数据库

	res := RespModel{
		StatusCode: 0,
		StatusMsg:  "success",
	}
	c.JSON(200, res)
}
