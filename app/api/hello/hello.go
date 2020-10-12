package hello

import (
    "github.com/gogf/gf/net/ghttp"
)

// Any output "Hello World!".
func Hello(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
}
