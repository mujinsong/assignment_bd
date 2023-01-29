package backend

import (
	"assignment_bd/dao"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []dao.Video `json:"video_list,omitempty"`
	NextTime  time.Time   `json:"next_time,omitempty"`
}
