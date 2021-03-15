package v1

import (
	"gf-music/library/utils"
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	utils.WeChatNotification("a", "b")

}
