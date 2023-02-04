package global

import (
	"github.com/hertz-contrib/jwt"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	HzJwtMw *jwt.HertzJWTMiddleware
)

