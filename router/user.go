package router

import (
	v1 "gf-music/app/api/v1"
	"gf-music/app/middleware"
	"github.com/gogf/gf/frame/g"
)

// InitAdminsRouter 注册管理员路由
func InitUsersRouter() {
	UserRouter := g.Server().Group("user").Middleware(
		middleware.JwtAuth,
		middleware.CORS,
	)
	{
		UserRouter.POST("hello", v1.Hello)
		UserRouter.POST("login", v1.GfJWTMiddleware.LoginHandler)
	}
}
