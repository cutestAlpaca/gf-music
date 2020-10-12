package router

import (
    "gf-music/app/api/hello"
	"gf-music/app/api/web/user"
	"gf-music/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.SetRewrite("/favicon.ico", "response/image/favicon.ico")

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", hello.Hello)
	})

	api := s.Group("/api")
	api.Middleware(middleware.CORS)
	api.Middleware(middleware.Auth)
	api.ALL("/", func(r *ghttp.Request) {
		_ = r.Response.WriteJson(map[string]interface{}{
			"Code": 200,
			"msg":  "",
			"data": "api success",
		})
	})
	ctlUser := new(user.Controller)
	api.ALL("/user", ctlUser)

	//api.Group("/user", func(group *ghttp.RouterGroup) {
	//	group.Middleware(middleware.Auth)
	//
	//})
}
