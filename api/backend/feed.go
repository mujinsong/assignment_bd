package backend

import (
	"assignment_bd/model"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  time.Time     `json:"next_time,omitempty"`
}
