package boot

import (
	"fmt"
	"gf-music/library/service"
	_ "gf-music/packed"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	_ = gtime.SetTimeZone("Asia/Shanghai") //设置系统时区

	//initGfToken()

	showLogo()
}

func initGfToken()  {
	// 启动 gtoken
	gtoken := &gtoken.GfToken{
		CacheMode:        g.Cfg().GetInt8("gToken.CacheMode"),
		CacheKey:         g.Cfg().GetString("gToken.CacheKey"),
		Timeout:          g.Cfg().GetInt("gToken.Timeout"),
		MaxRefresh:       g.Cfg().GetInt("gToken.MaxRefresh"),
		TokenDelimiter:   g.Cfg().GetString("gToken.TokenDelimiter"),
		EncryptKey:       g.Cfg().GetBytes("gToken.EncryptKey"),
		AuthFailMsg:      g.Cfg().GetString("gToken.AuthFailMsg"),
		LoginPath:        "/api/login",
		LoginBeforeFunc:  service.AdminLogin,
		LoginAfterFunc:   service.LoginAfter,
		LogoutPath:       "/api/logout",
		AuthPaths:        g.SliceStr{"/api/*"},
		AuthAfterFunc:    service.AuthAfterFunc,
		LogoutBeforeFunc: service.LoginOut,                     // 开启全局拦截，默认关闭
	}
	gtoken.Start()
}

func showLogo() {
	fmt.Println("----------------")
	fmt.Println("当前版本:" + service.Version)
}