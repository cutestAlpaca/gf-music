package response

import (
	"gf-music/app/model/user"
	"github.com/gogf/gf/os/gtime"
)

// AdminLogin response Structure
type Login struct {
	User      *user.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

// AdminResponse response Structure
type Response struct {
	Admin *user.Entity `json:"user"`
}

type LoginFormat struct {
	Id         int         `json:"id"`          // 主键
	Uuid       string      `json:"uuid"`        // UUID
	Username   string      `json:"username"`    // 登录名
	Sex        int         `json:"sex"`         // 性别;0:保密,1:男,2:女
	Enable     int         `json:"enable"`      // 是否启用//radio/1,启用,2,禁用
	UpdateTime *gtime.Time `json:"update_time"` // 更新时间
	CreateTime *gtime.Time `json:"create_time"` // 创建时间
	IsAdmin    int         `json:"is_admin"`    // 是否后台管理员 1 是  0   否
	Remark     string      `json:"remark"`      // 备注
	Avatar     string      `json:"avatar"`      // 头像
	UserEmail  string      `json:"user_email"`  // 用户邮箱
}
