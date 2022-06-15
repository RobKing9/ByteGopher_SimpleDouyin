package controller

import (
	"ByteGopher_SimpleDouyin/dao"
	"ByteGopher_SimpleDouyin/model"
	"ByteGopher_SimpleDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type FavoriteController interface {
	FavoriteAction(c *gin.Context)
	GetFavouriteList(c *gin.Context)
}
type favoriteController struct {
	favoriteDao dao.FavoriteDao
}

func NewFavoriteController() FavoriteController {
	return &favoriteController{
		favoriteDao: dao.NewFavoriteDao(),
	}
}

type likeResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type GetFavouriteListResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	VideoList  []model.Video `json:"video_list,omitempty"`
}

// FavoriteAction 点赞或者取消赞操作;
func (fc *favoriteController) FavoriteAction(c *gin.Context) {
	// 若鉴权失败，则返回
	flag, exist := c.Get("flag")
	fmt.Println("鉴权flag", flag)
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: this flag not exist")
	}
	if !flag.(bool) { // 认证失败
		c.JSON(http.StatusOK, RelationActionResponse{
			RespModel: model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg:  "no permission",
			},
		})
		log.Println("请先登录！！！！！")
		return
	}
	strUserId, _ := c.Get("userid")
	fmt.Println("获取到的用户id：", strUserId.(int64))
	strVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	strActionType := c.Query("action_type")
	actionType, _ := strconv.ParseInt(strActionType, 10, 64)
	//获取点赞或者取消赞操作的错误信息
	err := fc.favoriteDao.UpdateLike(strUserId.(int64), videoId, int32(actionType))
	if err == nil {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 成功")
		c.JSON(http.StatusOK, likeResponse{
			StatusCode: 0,
			StatusMsg:  "favourite action success",
		})
	} else {
		log.Printf("方法like.FavouriteAction(userid, videoId, int32(actiontype) 失败：%v", err)
		c.JSON(http.StatusOK, likeResponse{
			StatusCode: 1,
			StatusMsg:  "favourite action fail",
		})
	}
}

// GetFavouriteList 获取点赞列表;
func (fc *favoriteController) GetFavouriteList(c *gin.Context) {
	// 若鉴权失败，则返回
	flag, exist := c.Get("flag")
	if !exist {
		// TODO: 存入日志？
		log.Println("RelationAction: this flag not exist")
	}
	if !flag.(bool) { // 认证失败
		c.JSON(http.StatusOK, RelationActionResponse{
			RespModel: model.RespModel{
				StatusCode: model.SCodeFalse,
				StatusMsg:  "no permission",
			},
		})
		return
	}
	strUserId := c.Query("user_id")
	//strCurId := c.GetString("userId")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	//curId, _ := strconv.ParseInt(strCurId, 10, 64)
	videoIds, err := fc.favoriteDao.GetLikeUserIdList(userId)
	if err != nil {
		log.Printf("方法fc.favoriteDao.GetLikeUserIdList(userId) 失败：%v", err)
		c.JSON(http.StatusOK, GetFavouriteListResponse{
			StatusCode: 1,
			StatusMsg:  "get favouriteList fail ",
		})
		return
	}
	var wg sync.WaitGroup
	var videoModels []model.VideoModel
	for i, _ := range videoIds {
		wg.Add(1)
		go func() {
			videoModel, err := fc.favoriteDao.GetVideoModelByID(i)
			if err != nil {
				log.Printf("方法fc.favoriteDao.GetVideoModelByID(i) 失败：%v", err)
				return
			}
			videoModels = append(videoModels, *videoModel)
			defer wg.Done()
		}()
	}
	wg.Wait()
	videos := utils.ConvertVideoModelListToVideoList(videoModels)
	for i, _ := range videos {
		videos[i].IsFavorite = true
	}
	if err == nil {
		log.Printf("方法like.GetFavouriteList(userid) 成功")
		c.JSON(http.StatusOK, GetFavouriteListResponse{
			StatusCode: 0,
			StatusMsg:  "get favouriteList success",
			VideoList:  videos,
		})
	} else {
		log.Printf("方法like.GetFavouriteList(userid) 失败：%v", err)
		c.JSON(http.StatusOK, GetFavouriteListResponse{
			StatusCode: 1,
			StatusMsg:  "get favouriteList fail ",
		})
	}
}
