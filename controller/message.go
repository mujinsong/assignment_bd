package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
)

// MessageChat 是获取聊天记录
// 因为客户端会定时轮询此接口来查询消息记录，所以每次只读取未读过的消息
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var messages []model.Message
	statusCode := consts.STATUS_SUCCESS
	statusMsg := "获取聊天记录成功"
	ownID, _ := utils.GetUid(c)

	// service 进行逻辑处理
	messages, err = service.GetMessageList(ownID, utils.StrToUint64(c.Query("to_user_id")))

	// 捕捉异常并返回给客户端
	if err != nil {
		log.Println(err)
		statusCode = consts.STATUS_FAILURE
		statusMsg = "获取聊天记录失败"
	}

	// 向客户端返回响应
	c.JSON(http.StatusOK, model.MessageResponse{
		Response: model.Response{
			StatusCode: statusCode,
			StatusMsg:  statusMsg,
		},
		MessageList: messages,
	})
}

// MessageAction 是发送消息操作
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	statusCode := consts.STATUS_SUCCESS
	statusMsg := "信息发送成功"
	fromUserID, _ := utils.GetUid(c)
	toUserID := utils.StrToUint64(c.Query("to_user_id"))
	content := c.Query("content")

	err = service.SendMessage(fromUserID, toUserID, content)

	if err != nil {
		statusCode = consts.STATUS_SUCCESS
		statusMsg = "信息发送失败"
		log.Println(err)
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: statusCode,
		StatusMsg:  statusMsg,
	})
}
