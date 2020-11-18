package v1

import (
	"github.com/gogf/gf/net/ghttp"
)

func Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}
