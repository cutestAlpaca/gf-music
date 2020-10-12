package user

import (
	"errors"
	"fmt"
	"gf-music/app/model/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

const (
	USER_SESSION_MARK = "user_info"
)

// 注册输入参数
type SignUpInput struct {
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Email  string `v:"required"`
	Mobile  string `v:"required|length:1,11#手机不能为空|账号长度应当在:min到:max之间"`
	Nickname  string
	IpAddress string
}

// 用户注册
func SignUp(data *SignUpInput) error {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.FirstString())
	}
	// 昵称为非必需参数，默认使用账号名称
	if data.Nickname == "" {
		data.Nickname = data.Email
	} else if data.Email == "" {
		data.Nickname = data.Mobile
	}
	// 手机号唯一性数据检查
	if !CheckMobile(data.Mobile) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Mobile))
	}
	// 邮箱唯一性数据检查
	if !CheckEmail(data.Email) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Email))
	}
	// 昵称唯一性数据检查
	if !CheckNickName(data.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data.Nickname))
	}
	// 将输入参数赋值到数据库实体对象上
	var entity *user.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}
	if _, err := user.Save(entity); err != nil {
		return err
	}
	return nil
}

// 判断用户是否已经登录
func IsSignedIn(session *ghttp.Session) bool {
	return session.Contains(USER_SESSION_MARK)
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func SignIn(passport, password string, session *ghttp.Session) error {
	one, err := user.FindOne("mobile=? and password=?", passport, password)
	if err != nil {
		return err
	}
	if one == nil {
		return errors.New("账号或密码错误")
	}
	return session.Set(USER_SESSION_MARK, one)
}

// 用户注销
func SignOut(session *ghttp.Session) error {
	return session.Remove(USER_SESSION_MARK)
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckMobile(mobile string) bool {
	if i, err := user.FindCount("mobile", mobile); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckEmail(email string) bool {
	if i, err := user.FindCount("email", email); err != nil {
		return false
	} else {
		return i == 0
	}
}


// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func CheckNickName(nickname string) bool {
	if i, err := user.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 获得用户信息详情
func GetProfile(session *ghttp.Session) (u *user.Entity) {
	_ = session.GetStruct(USER_SESSION_MARK, &u)
	return
}
