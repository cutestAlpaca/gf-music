package boot

import (
	"fmt"
	"gf-music/app/service"
	_ "gf-music/packed"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	_ = gtime.SetTimeZone("Asia/Shanghai") //设置系统时区

	showLogo()
}

func showLogo() {
	fmt.Println("----------------")
	fmt.Println("当前版本:" + service.Version)
}
