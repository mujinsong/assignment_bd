package controller

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
)

// Register 用户注册账号（剩下逻辑注释本方法作者补写）
func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	// 注册用户到数据库（service 层的操作）
	//println("注册操作")
	_, err := service.Register(username, password)
	if err != nil {
		c.JSON(consts.SUCCESS, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}

	token := jwt.New(jwt.GetSigningMethod(global.HzJwtMw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)
	id, err := service.GetUserIDByUsername(username)
	if err != nil {
		c.JSON(consts.SUCCESS, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}
	claims[consts.IdentityKey] = id

	expire := global.HzJwtMw.TimeFunc().Add(global.HzJwtMw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = global.HzJwtMw.TimeFunc().Unix()
	tokenString, err := token.SignedString(global.HzJwtMw.Key)
	if err != nil {
		return
	}

	global.HzJwtMw.LoginResponse(ctx, c, consts.SUCCESS, tokenString, expire)
	//global.HzJwtMw.LoginHandler(ctx, c)
}

// Login 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	_, err := service.Login(&model.Login{Username: username, Password: password})
	if err != nil {
		c.JSON(consts.SUCCESS, model.Response{StatusCode: consts.STATUS_FAILURE, StatusMsg: err.Error()})
		return
	}

	global.HzJwtMw.LoginHandler(ctx, c)
}

// UserInfo 获取用户信息（剩下逻辑注释本方法作者补写）
func UserInfo(ctx context.Context, c *app.RequestContext) {
	// 获取指定用户的 ID，并请求用户详细信息 UserInfo
	userID := utils.StrToUint64(c.Query("user_id"))
	uid, _ := utils.GetUid(c)
	userInfo := service.UserInfoGetByUserID(userID, uid)
	c.JSON(consts.SUCCESS, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: consts.STATUS_SUCCESS,
			StatusMsg:  "进入个人信息页面",
		},
		UserInfo: &userInfo,
	})
}
