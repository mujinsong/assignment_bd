package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// MessageChat 是聊天记录（剩下逻辑注释本方法作者补写）
func MessageChat(ctx context.Context, c *app.RequestContext) {

}

// MessageAction 是消息操作（剩下逻辑注释本方法作者补写）
func MessageAction(ctx context.Context, c *app.RequestContext) {
	c.URI().String()

}
