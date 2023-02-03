package controller

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// RelationAction 用户之间的关系操作（剩下逻辑注释本方法作者补写）
func RelationAction(ctx context.Context, c *app.RequestContext) {

}

// FollowList 用来获取用户关注列表（剩下逻辑注释本方法作者补写）
func FollowList(ctx context.Context, c *app.RequestContext) {

}

// FollowerList 用来获取粉丝列表（剩下逻辑注释本方法作者补写）
func FollowerList(ctx context.Context, c *app.RequestContext) {

}

// FriendList 获取用户好友列表（剩下逻辑注释本方法作者补写）
func FriendList(ctx context.Context, c *app.RequestContext) {
	var users []model.User
	global.DB.Where("user_id = ?", c.Query("user_id")).Find(&users)

	c.JSON(http.StatusOK, model.FriendListResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  "查询朋友列表成功",
		},
		FriendList: []model.FriendUser{
			model.FriendUser{
				UserInfo: model.UserInfo{
					ID:            1,
					Name:          "lipanke",
					FollowCount:   0,
					FollowerCount: 0,
					IsFollow:      false,
					Avatar:        "http://img.panker916.space/32bf3ac8c3e1420bafee3ab84cb5f17e",
				},
				Message: "这是一条测试消息",
				MsgType: 1,
			},
		},
	})
}
