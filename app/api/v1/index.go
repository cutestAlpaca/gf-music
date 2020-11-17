package v1

import (
	"gf-music/app/api/request"
	"gf-music/app/service"
	"gf-music/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Register(r *ghttp.Request) {
	var R *request.Register
	if err := r.Parse(&R); err != nil {
		global.FailWithMessage(r, err.Error())
		r.Exit()
	}
	if err := service.Register(R); err != nil {
		global.FailWithMessage(r, err.Error())
		r.ExitAll()
	}
	global.OkDetailed(r, g.Map{}, "注册成功!")
}
