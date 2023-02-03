package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// Publish 发布视频的操作 （剩下逻辑注释本方法作者补写）
func Publish(ctx context.Context, c *app.RequestContext) {

}

// PublishList 根据 user_id 查询用户 id，再查询这个用户发布的视频 （剩下逻辑注释本方法作者补写）
// TODO 这里的视频列表里的视频都是写死的，以后可以考虑用 oss 来存储，数据库里存储 URL
func PublishList(ctx context.Context, c *app.RequestContext) {

}
