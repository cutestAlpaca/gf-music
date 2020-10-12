package middleware

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Auth(r *ghttp.Request)  {
	// 启动gtoken
	gtoken := &gtoken.GfToken{
		LoginPath:       "/login",
		LoginBeforeFunc: Login,
		LogoutPath:      "/user/logout",
		AuthPaths:        g.SliceStr{"/api/user", "/system"}, // 这里是按照前缀拦截，拦截/user /user/list /user/add ...
		GlobalMiddleware: false,                           // 开启全局拦截，默认关闭
		AuthFailMsg: "请求错误或登录超时111111～！",
	}
	gtoken.Start()
}

func Login(r *ghttp.Request) (string ,interface{}) {
	//username := r.GetPostString("username")
	//passwd := r.GetPostString("passwd")

	// TODO 进行登录校验
	//r.Middleware.Next()

	return "yixiaohu", []g.MapStrStr{{"name": "张三", "age": "18"}, {"name": "李四", "age": "32"}}
}