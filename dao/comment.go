package dao

import (
	"ByteGopher_SimpleDouyin/model"
	"errors"
	"log"
)

type CommentDao interface {
	Count(videoId int64) (int64, error)
	CommentIdList(videoId int64) ([]string, error)
	InsertComment(comment model.CommentModel) (model.CommentModel, error)
	DeleteComment(commentId int64) error
	GetCommentList(videoId int64) ([]model.CommentModel, error)
}
type commentDao struct {
}

func NewCommentDao() CommentDao {
	return &commentDao{}
}

// Count
// 1、使用video id 查询Comment数量
func (com *commentDao) Count(videoId int64) (int64, error) {
	log.Println("CommentDao-Count: running") //函数已运行
	//Init()
	var count int64
	//数据库中查询评论数量
	err := MysqlDb.Model(model.CommentModel{}).Where("video_id = ?", videoId).Count(&count).Error
	if err != nil {
		log.Println("CommentDao-Count: return count failed") //函数返回提示错误信息
		return -1, errors.New("find comments count failed")
	}
	log.Println("CommentDao-Count: return count success") //函数执行成功，返回正确信息
	return count, nil
}

//CommentIdList 根据视频id获取评论id 列表
func (com *commentDao) CommentIdList(videoId int64) ([]string, error) {
	var commentIdList []string
	err := MysqlDb.Model(model.CommentModel{}).Select("id").Where("video_id = ?", videoId).Find(&commentIdList).Error
	if err != nil {
		log.Println("CommentIdList:", err)
		return nil, err
	}
	return commentIdList, nil
}

// InsertComment
// 2、发表评论
func (com *commentDao) InsertComment(comment model.CommentModel) (model.CommentModel, error) {
	log.Println("CommentDao-InsertComment: running") //函数已运行
	//数据库中插入一条评论信息
	log.Println("!!!!===", comment.CommentID)
	//err := MysqlDb.Model(model.CommentModel{}).Create(&comment).Error
	err := MysqlDb.Save(&comment).Error
	if err != nil {
		log.Println("CommentDao-InsertComment: return create comment failed") //函数返回提示错误信息
		return model.CommentModel{}, errors.New("create comment failed")
	}
	log.Println("CommentDao-InsertComment: return success") //函数执行成功，返回正确信息
	return comment, nil
}

// DeleteComment
// 3、删除评论，传入评论id
func (com *commentDao) DeleteComment(commentId int64) error {
	log.Println("CommentDao-DeleteComment: running") //函数已运行
	var commentInfo model.CommentModel
	//先查询是否有此评论
	result := MysqlDb.Model(model.CommentModel{}).Where("comment_id = ?", commentId).First(&commentInfo)
	if result.RowsAffected == 0 { //查询到此评论数量为0则返回无此评论
		log.Println("CommentDao-DeleteComment: return del comment is not exist") //函数返回提示错误信息
		return errors.New("del comment is not exist")
	}
	//数据库中删除评论-更新评论状态为-1
	err := MysqlDb.Model(model.CommentModel{}).Where("comment_id = ?", commentId).Delete(&model.CommentModel{}).Error
	if err != nil {
		log.Println("CommentDao-DeleteComment: return del comment failed") //函数返回提示错误信息
		return errors.New("del comment failed")
	}
	log.Println("CommentDao-DeleteComment: return success") //函数执行成功，返回正确信息
	return nil
}

// GetCommentList
// 4.根据视频id查询所属评论全部列表信息
func (com *commentDao) GetCommentList(videoId int64) ([]model.CommentModel, error) {
	log.Println("CommentDao-GetCommentList: running") //函数已运行
	//数据库中查询评论信息list
	var commentList []model.CommentModel
	result := MysqlDb.Model(model.CommentModel{}).Where("video_id=?", videoId).Order("create_date desc").Find(&commentList)
	//若此视频没有评论信息
	if result.RowsAffected == 0 {
		log.Println("CommentDao-GetCommentList: return there are no comments") //函数返回提示无评论
		return commentList, errors.New("there are no comments")
	}
	//若获取评论列表出错
	if result.Error != nil {
		log.Println(result.Error.Error())
		log.Println("CommentDao-GetCommentList: return get comment list failed") //函数返回提示获取评论错误
		return commentList, errors.New("get comment list failed")
	}
	log.Println("CommentDao-GetCommentList: return commentList success") //函数执行成功，返回正确信息
	return commentList, nil
}
