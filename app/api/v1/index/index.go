package v1

import (
	"gf-music/app/api/request"
	"gf-music/library/response"
	"github.com/gogf/gf/net/ghttp"
)

func Register(r *ghttp.Request) {
	var R *request.Register
	if err := r.Parse(&R); err != nil {
		response.FailWithMessage(r, err.Error())
		r.Exit()
	}

}
