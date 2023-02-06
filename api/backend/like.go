package backend

import (
	"assignment_bd/dao"
)

type Video struct {
	Id            uint   `json:"id,omitempty"`
	Author        User   `json:"author,omitempty"`
	Title         string `json:"title,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type LikeList struct {
	Response
	VideoList []dao.Video `json:video_list`
}
