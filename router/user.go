package router

import (
	v1 "gf-music/app/api/v1"
	"gf-music/app/middleware"
	"github.com/gogf/gf/frame/g"
)

// InitAdminsRouter 注册管理员路由
func InitUsersRouter() {
	UserRouter := g.Server().Group("api/user").Middleware(
		middleware.CORS,
		middleware.JwtAuth,
	)
	{
<<<<<<< HEAD
		UserRouter.ALL("hello", v1.Hello)
=======
		UserRouter.POST("hello", v1.Hello)
>>>>>>> 9fe917b75b6123615d240c665a2d6bcb3a47d694
	}
}
