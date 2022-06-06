package model

type Video struct {
	// 视频唯一标识
	Id int64 `json:"id" gorm:"column:video_id"`
	// 视频作者信息
	Author User `json:"author" gorm:"embedded"`
	// 视频播放地址
	PlayURL string `json:"play_url" gorm:"column:play_url"`
	// 视频封面地址
	CoverURL string `json:"cover_url" gorm:"column:cover_url"`
	// 视频的点赞总数
	FavoriteCount int64 `json:"favorite_count" gorm:"column:favorite_count"`
	// 视频的评论总数
	CommentCount int64 `json:"comment_count" gorm:"column:comment_count"`
	// 是否点赞
	// true-已点赞，false-为点赞
	IsFavorite bool `json:"is_favorite" gorm:"column:is_favorite"`
	// 视频标题
	Title string `json:"title" gorm:"column:title"`
}
