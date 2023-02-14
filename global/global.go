package global

import (
	"github.com/hertz-contrib/jwt"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	HzJwtMw *jwt.HertzJWTMiddleware
	// 时间游标，string是信息的联合id，uint64是上次读过的unix时间值
)
