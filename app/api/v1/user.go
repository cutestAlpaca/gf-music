package v1

import (
	"github.com/gogf/gf/net/ghttp"
	"time"
)

func Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}

func Login(r *ghttp.Request)  {
	username := r.GetParam("username")
	password := r.GetParam("password")
	if username == "admin" && password == "123" {
		LoginResponse(r,200,"aaa",time.Now())
	}
}