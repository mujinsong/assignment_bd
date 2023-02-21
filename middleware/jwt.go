package middleware

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

func MyJwt() {
	var err error
	global.HzJwtMw, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte("tiktok jwt key"),
		//Timeout:       time.Hour,
		MaxRefresh:  0,                           // token 永不过期 方便测试
		TokenLookup: "query: token, form: token", //"header: Authorization", //  , cookie: jwt query: token ,
		//TokenHeadName: "Bearer",
		//Token 的返回
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			parseToken, err := global.HzJwtMw.ParseTokenString(token)
			if err != nil {
				return
			}
			claim := jwt.ExtractClaimsFromToken(parseToken)
			id := claim[consts.IdentityKey]
			c.JSON(consts.SUCCESS, model.UserLoginResponse{
				Response: model.Response{
					StatusCode: consts.STATUS_SUCCESS,
					StatusMsg:  "success",
				},
				UserID: uint64(id.(float64)),
				Token:  token,
			})
		},
		//用户登陆（认证）
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			user, err := service.Login(&model.Login{Username: loginStruct.Username, Password: loginStruct.Password})
			if err != nil {
				return nil, err
			}
			return user, nil
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err = %+v", e.Error())
			return e.Error()
		},
		//用户信息的提取
		IdentityKey: consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {

			claims := jwt.ExtractClaims(ctx, c)
			//fmt.Println("claims:", claims)
			//fmt.Println("nothing else matter", int(claims[IdentityKey].(float64)))
			return model.User{
				ID: uint64(int(claims[consts.IdentityKey].(float64))),
			}
		},
		//通过授权者
		//Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
		//	//id, err := strconv.Atoi(c.Query("id"))
		//	//if err != nil {
		//	//	return false
		//	//}
		//	//if v, ok := data.(*model.User); ok && v.ID == uint(id) {
		//	//	return true
		//	//}
		//	//return false
		//	return true
		//},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(consts.SUCCESS, model.Response{
				StatusCode: 1,
				StatusMsg:  message,
			})
		},
		//todo 把负载
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
	})
	if err != nil {
		panic(err)
	}
}

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get("username")
	c.JSON(200, utils.H{
		"message": fmt.Sprintf("id:%v", user.(*model.User).ID),
	})
}
