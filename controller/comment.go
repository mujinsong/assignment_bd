package controller

import (
	"assignment_bd/consts"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// CommentAction 用户之间进行聊天操作（剩下逻辑注释本方法作者补写）
func CommentAction(ctx context.Context, c *app.RequestContext) {
	// 解析参数
	var comment model.Comments

	payload, exist := c.Get(consts.IdentityKey)
	if !exist {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	uid := payload.(model.User).ID
	comment.UserID = uid

	comment.VideoID = utils.StrToUint64(c.Query("video_id"))
	comment.ActionType = utils.StrToUint8(c.Query("action_type"))
	comment.Content = c.Query("comment_text")
	comment.CreateDate = utils.CurrentTime()

	// 修改视频的评论数
	service.UpdateVideoCommentCount(comment.VideoID, comment.ActionType)

	// 修改评论信息
	if c.Query("action_type") == "1" {
		// 发表评论
		err := service.CreateComment(&comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Response{
				StatusCode: consts.STATUS_FAILURE,
				StatusMsg:  "操作失败"})
			return
		}
		// 返回评论ID
		////println("成功发表评论,用户ID为", comment.UserID)

		c.JSON(consts.SUCCESS, model.CommentResponse{
			Response: model.Response{
				StatusCode: consts.STATUS_SUCCESS,
				StatusMsg:  "success"},
			Comment: model.Comment{
				ID: comment.ID,
				// 这里返回的是自己的信息
				User:       service.UserInfoGetByUserID(comment.UserID, comment.UserID),
				Content:    comment.Content,
				CreateDate: comment.CreateDate,
			},
		})
	} else {
		// 删除评论
		// 默认 comment.ActionType == 2
		err := service.DeleteComment(utils.StrToUint64(c.Query("comment_id")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Response{
				StatusCode: consts.STATUS_FAILURE,
				StatusMsg:  consts.MsgFlags[consts.ERROR]})
			return
		}
		c.JSON(consts.SUCCESS, model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  consts.MsgFlags[consts.SUCCESS],
		})
	}
}

// CommentList 获取评论列表
func CommentList(ctx context.Context, c *app.RequestContext) {
	////println("获取评论列表")
	// 解析参数
	uid, _ := utils.GetUid(c)
	videoID := utils.StrToUint64(c.Query("video_id"))
	// 获取评论列表
	comments, err := service.GetCommentList(videoID, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: consts.STATUS_FAILURE,
			StatusMsg:  "操作失败"})
		return
	}
	// 返回评论列表
	c.JSON(consts.SUCCESS, model.CommentListResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  "success"},
		CommentList: comments,
	})
}
