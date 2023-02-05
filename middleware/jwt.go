package middleware

import (
	"assignment_bd/consts"
	"assignment_bd/global"
	"assignment_bd/model"
	"assignment_bd/service"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

func MyJwt() {
	var err error
	global.HzJwtMw, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte("tiktok jwt key"),
		//Timeout:       time.Hour,
		MaxRefresh:  time.Hour,
		TokenLookup: "query: token, form: token", //"header: Authorization", //  , cookie: jwt query: token ,
		//TokenHeadName: "Bearer",
		//Token 的返回
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
				"message": "success",
			})
			//c.JSON(http.StatusOK, backend.UserLoginResponse{
			//	Response: backend.Response{StatusCode: 200, StatusMsg: "OK"},
			//})
		},
		//用户登陆（认证）
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			fmt.Println("problem in this")
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			user, err := service.Login(&model.Login{Username: loginStruct.Username, Password: loginStruct.Password})
			if err != nil {
				return nil, err
			}
			fmt.Println("user:", user)
			//c.JSON(http.StatusOK, backend.UserLoginResponse{
			//	Response: backend.Response{StatusCode: 200, StatusMsg: "OK"},
			//	UserID:   uint64(user.Id),
			//	Token:
			//})
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
				Id: int64(int(claims[consts.IdentityKey].(float64))),
			}
		},
		//通过授权者
		//Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
		//	//id, err := strconv.Atoi(c.Query("id"))
		//	//if err != nil {
		//	//	return false
		//	//}
		//	//if v, ok := data.(*model.User); ok && v.Id == uint(id) {
		//	//	return true
		//	//}
		//	//return false
		//	return true
		//},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
		//todo 把负载
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.Id,
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
		"message": fmt.Sprintf("id:%v", user.(*model.User).Id),
	})
}
