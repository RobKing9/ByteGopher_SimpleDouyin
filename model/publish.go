package model

type VideoModel struct {
	Id            int64
	Author        User
	PlayUrl       string
	CoverUrl      string // 封面的url
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
	Title         string
}
