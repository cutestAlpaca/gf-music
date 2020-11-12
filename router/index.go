package router

import (
	v1 "gf-music/app/api/v1"
	"gf-music/app/middleware"
	"github.com/gogf/gf/frame/g"
)

func InitIndexRouter()  {
	IndexRouter := g.Server().Group("index").Middleware(
		middleware.CORS,
	)
	{
		IndexRouter.POST("register",v1.GfJWTMiddleware.LoginHandler)
		IndexRouter.POST("login",v1.GfJWTMiddleware.LoginHandler)
		IndexRouter.POST("refresh", v1.GfJWTMiddleware.RefreshHandler)
	}
}
