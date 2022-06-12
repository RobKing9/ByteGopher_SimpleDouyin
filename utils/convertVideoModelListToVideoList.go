package utils

import (
	"ByteGopher_SimpleDouyin/model"
)

func ConvertVideoModelListToVideoList(vml []model.VideoModel) []model.Video {
	Videos := []model.Video{}
	for _, v := range vml {
		video := model.Video{
			Id: v.VideoID,
			Author: model.User{
				Id:            v.Author.UserID,
				Name:          v.Author.UserName,
				FollowCount:   v.Author.FollowCount,
				FollowerCount: v.Author.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
		Videos = append(Videos, video)
	}
	return Videos
}

////curId表示当前登录用户id,通过curId判断是否点赞，关注等
//func ConvertVideoModelListToVideoListLogin(vml []model.VideoModel,curId int64) ([]model.Video,error) {
//	Videos := []model.Video{}
//	for _, v := range vml {
//		GivelikeInfo,err:=dao.NewFavoriteDao().GetLikeInfo(curId,v.VideoID)
//		if err!=nil{
//			return nil, err
//		}
//		isFavorite:=GivelikeInfo.Status
//		video := model.Video{
//			Id:            v.VideoID,
//			Author:        model.User{
//				Id:            v.Author.UserID,
//				Name:          v.Author.UserName,
//				FollowCount:   v.Author.FollowCount,
//				FollowerCount: v.Author.FollowerCount,
//				//关注接口写了以后，可以添加判断是否关注
//				IsFollow:     false,
//			},
//			PlayUrl:       v.PlayURL,
//			CoverUrl:      v.CoverURL,
//			FavoriteCount: v.FavoriteCount,
//			CommentCount:  v.CommentCount,
//			IsFavorite:    isFavorite,
//			Title:         v.Title,
//		}
//		Videos=append(Videos,video)
//	}
//	return Videos,nil
//}
