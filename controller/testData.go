package controller

import (
	"tiktok/config"
	"time"
)

var (
	TestUser = User{
		Id:            1,
		Name:          "测试用户",
		FollowCount:   114514,
		FollowerCount: 1919810,
		IsFollow:      false,
	}
	TestUserInfoResponse = UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		User:       TestUser,
	}

	TestVideoList = []Video{
		{
			Id:            1,
			Author:        TestUser,
			PlayUrl:       config.DefaultHost+ "public/test1.mp4",
			CoverUrl:      "https://i0.hdslb.com/bfs/article/a2617bd0ae3fa13281af1fff146ce795e3c45211.jpg@942w_1424h_progressive.webp",
			FavoriteCount: 114514,
			CommentCount:  1919810,
			IsFavorite:    false,
		},
		{
			Id:            2,
			Author:        TestUser,
			PlayUrl:       config.DefaultHost + "public/test2.mp4",
			CoverUrl:      "https://i0.hdslb.com/bfs/article/54fff10abcedf6eeae791214add1abd67e299d38.jpg@942w_1413h_progressive.webp",
			FavoriteCount: 114514,
			CommentCount:  1919810,
			IsFavorite:    false,
		},
	}

	TestFeedResponse = FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频列表成功",
		VideoList:  TestVideoList,
		NextTime:   time.Now().Unix(),
	}
)
