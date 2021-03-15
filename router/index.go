package router

import (
	v1 "gf-music/app/api/v1"
	"gf-music/app/middleware"
	"github.com/gogf/gf/frame/g"
)

func InitIndexRouter() {
	IndexRouter := g.Server().Group("api/index").Middleware(
		middleware.CORS,
	)
	{
		//test api
		IndexRouter.ALL("hello", v1.Hello)

		//
		IndexRouter.POST("register", v1.Register)
		IndexRouter.POST("login", v1.GfJWTMiddleware.LoginHandler)
		IndexRouter.POST("index", v1.Index)
		IndexRouter.POST("refresh", v1.GfJWTMiddleware.RefreshHandler)
		IndexRouter.POST("login", v1.GfJWTMiddleware.LoginHandler)
	}
}
