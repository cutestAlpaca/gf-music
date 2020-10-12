package user

import (
	"gf-music/app/service/user"
	"gf-music/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}

// 注册请求参数，用于前后端交互参数格式约定
type SignUpRequest struct {
	user.SignUpInput
}

// 用户注册接口
func (c *Controller) SignUp(r *ghttp.Request) {
	var data *SignUpRequest
	// 这里没有使用Parse而是仅用GetStruct获取对象，
	// 数据校验交给后续的service层统一处理
	if err := r.GetStruct(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignUp(&data.SignUpInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		data.IpAddress = r.GetClientIp()
		response.JsonExit(r, 0, "ok", map[string] interface{} {
			"data" :data,
			"ip": data.IpAddress,
		})
	}
}

// 登录请求参数，用于前后端交互参数格式约定
type SignInRequest struct {
	Mobile string `v:"required#手机号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// 用户登录接口
func (c *Controller) SignIn(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := user.SignIn(data.Mobile, data.Password, r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

// 判断用户是否已经登录
func (c *Controller) IsSignedIn(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.IsSignedIn(r.Session))
}

// 用户注销/退出接口
func (c *Controller) SignOut(r *ghttp.Request) {
	if err := user.SignOut(r.Session); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckPassportRequest struct {
	Mobile string
}


func (c *Controller) CheckPassport(r *ghttp.Request) {
	var data *CheckPassportRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Mobile != "" && !user.CheckMobile(data.Mobile) {
		response.JsonExit(r, 0, "账号已经存在", false)
	}
	response.JsonExit(r, 0, "", true)
}

// 账号唯一性检测请求参数，用于前后端交互参数格式约定
type CheckNickNameRequest struct {
	Nickname string
}


func (c *Controller) CheckNickName(r *ghttp.Request) {
	var data *CheckNickNameRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if data.Nickname != "" && !user.CheckNickName(data.Nickname) {
		response.JsonExit(r, 0, "昵称已经存在", false)
	}
	response.JsonExit(r, 0, "ok", true)
}

/**

 */
func (c *Controller) Profile(r *ghttp.Request) {
	response.JsonExit(r, 0, "", user.GetProfile(r.Session))
}
