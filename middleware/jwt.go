package middleware

import (
	"assignment_bd/consts"
	"assignment_bd/dao"
	"assignment_bd/global"
	"assignment_bd/service"
	"assignment_bd/utils"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hutils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"

	"log"
	"net/http"
	"time"
)

// biz/router/middleware/jwt.go

// Claim 定义用户登陆信息结构体
type Claim struct {
	ID       int
	Username string
}

var IdentityKey string = "id"

func JwtMwInit() {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		// 置所属领域名称
		Realm: "hertz jwt",
		// 用于设置签名密钥
		Key: []byte(utils.RandStr(consts.TOKEN_LENGTH)),
		// 设置 token 过期时间
		Timeout: time.Hour * 8,
		// 设置最大 token 刷新时间
		MaxRefresh: time.Hour * 4,
		// 设置 token 的获取源
		TokenLookup: "header: Authorization",
		// 设置从 header 中获取 token 时的前缀
		TokenHeadName: "Bearer",
		// 用于设置检索身份的键
		IdentityKey: IdentityKey,
		// 登陆成功后为向 token 中添加自定义负载信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(Claim); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 从 token 提取用户信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			claim := claims[IdentityKey].(map[string]interface{})
			return &Claim{
				ID:       int(claim["ID"].(float64)),
				Username: claim["Username"].(string),
			}
		},
		// 认证
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginUser dao.Login
			if err := c.BindAndValidate(&loginUser); err != nil {
				return "", errors.New("数据绑定错误")
			}
			logined, err := service.Login(&loginUser)
			if err != nil {
				return nil, err
			}
			return logined, nil
		},
		// 鉴权
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			// 单一角色 不设权限校验
			return true
		},
		// 登录响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, hutils.H{
				"code":   code,
				"token":  token,
				"detail": "登陆成功",
			})
		},
		// 设置 jwt 校验流程发生错误时响应所包含的错误信息
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		// 无权限
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, hutils.H{
				"code":   10011,
				"detail": "无权限或用户认证已过期",
				"data":   nil,
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	global.HzJwtMw = authMiddleware
}
