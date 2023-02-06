package model

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"mime/multipart"
	"time"
)

// Video 对应的是数据库中的 video 结构，用来存储
type Video struct {
	Id            int64     `json:"id,omitempty"`
	UserId        int64     `json:"user_id,omitempty"`
	Title         string    `json:"title,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
}

// VideoInfo 视频信息，基本包含了视频的所有信息，不能直接通过数据库获取，需要各个字段拼装获取
type VideoInfo struct {
	ID            int64    `json:"id"`             // 视频唯一标识
	Author        UserInfo `json:"author"`         // 视频作者信息
	PlayURL       string   `json:"play_url"`       // 视频播放地址
	CoverURL      string   `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64    `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64    `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool     `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string   `json:"title"`          // 视频标题
}

// Publish 上传视频至数据库
func (i VideoInfo) Publish(data *multipart.FileHeader, userId int64, title string) error {
	// 视频播放地址
	file, err := data.Open()
	if err != nil {
		log.Printf("方法data.Open() 失败%v", err)
		return err
	}
	log.Printf("方法data.Open() 成功")

	//生成一个uuid作为视频的名字
	videoName := uuid.NewV4().String()
	log.Printf("生成视频名称%v", videoName)
	// 上传视频
	err = VideoFTP(file, videoName)
	if err != nil {
		log.Printf("方法model.VideoFTP(file, videoName) 失败%v", err)
		return err
	}
	log.Printf("方法model.VideoFTP(file, videoName) 成功")
	defer file.Close()

	// 视频封面地址
	imageName := uuid.NewV4().String()
	//向队列中添加消息
	//ffmpeg.Ffchan <- ffmpeg.Ffmsg{
	//	videoName,
	//	imageName,
	//}

	// 补充组装
	err = Save(videoName, imageName, userId, title)
	if err != nil {
		log.Printf("方法model.Save(videoName, imageName, userId) 失败%v", err)
		return err
	}
	log.Printf("方法model.Save(videoName, imageName, userId) 成功")
	return nil
}

// VideoFTP 上传视频 待补充
func VideoFTP(file multipart.File, name string) error {
	//err := ftp.MyFTP.Cwd("video")
	//if err != nil {
	//	log.Println("转到路径video失败！！！")
	//} else {
	//	log.Println("转到路径video成功！！！")
	//}
	//err = ftp.MyFTP.Stor(videoName+".mp4", file)
	//if err != nil {
	//	log.Println("上传视频失败！！！！！")
	//	return err
	//}
	//log.Println("上传视频成功！！！！！")
	return nil
}

// Save 补充组装 待补充
func Save(videoName string, imageName string, userId int64, title string) error {
	return nil
}

// VideoCommentCount 评论数
type VideoCommentCount struct {
	VideoId int
	Count   int
}

func (receiver Video) TableName() string {
	return "videos"
}
