package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

/*
	项目本该部署到服务器上利用nginx提供视频流，这里为了方便测试，
*/
//func VideoStream(ctx context.Context, c *app.RequestContext){
//	// 从请求中获取视频id
//
//}

// Feed 视频流接口（给客户端推送短视频）（剩下逻辑注释本方法作者补写）
// TODO 完善视频推送的逻辑 和 查询优化
func Feed(ctx context.Context, c *app.RequestContext) {
	// 查询数据库中的视频信息
	videoList := service.
}
