package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// RelationAction 用户之间的 关注或者取消关注的 关系操作
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	followerID, err := utils.GetUid(c)
	userID := c.Query("to_user_id")      // 对方用户的id
	actionType := c.Query("action_type") // 1是关注; 2是取消关注

	statusCode := consts.STATUS_SUCCESS
	var statusMsg string

	// 根据不同的操作类型来进行不同的操作（关注 || 取关 || 异常）
	switch actionType {
	case "1": // 进行关注操作
		err = service.Follow(utils.StrToUint64(userID), followerID)
		statusMsg = "关注成功"
	case "2": // 进行取关操作
		err = service.UnFollow(utils.StrToUint64(userID), followerID)
		statusMsg = "取消关注成功"
	default: // 客户端返回了错误的操作类型，抛出异常
		err = errors.New("错误的操作类型")
	}

	// 异常处理
	if err != nil {
		statusCode = consts.STATUS_FAILURE
		statusMsg = "操作失败"
		fmt.Println(err)
	}

	// 返回响应
	c.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: statusCode,
			StatusMsg:  statusMsg,
		},
	})

}

// FollowList 用来获取用户关注列表
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	statusCode := consts.STATUS_SUCCESS
	var statusMsg string
	userID := utils.StrToUint64(c.Query("user_id"))
	users, err := service.GetFollowList(userID) // 获取用户关注列表 service 层

	if err != nil {
		statusCode = consts.STATUS_FAILURE
		statusMsg = "获取失败"
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{
			StatusCode: statusCode,
			StatusMsg:  statusMsg,
		},
		UserList: users,
	})
}

// FollowerList 用来获取粉丝列表
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	statusCode := consts.STATUS_SUCCESS
	var statusMsg string
	userID := utils.StrToUint64(c.Query("user_id"))
	users, err := service.GetFollowerList(userID) // 获取用户关注列表 service 层

	if err != nil {
		statusCode = consts.STATUS_FAILURE
		statusMsg = "获取失败"
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{
			StatusCode: statusCode,
			StatusMsg:  statusMsg,
		},
		UserList: users,
	})
}

// FriendList 获取用户好友列表（剩下逻辑注释本方法作者补写）
func FriendList(ctx context.Context, c *app.RequestContext) {
}
