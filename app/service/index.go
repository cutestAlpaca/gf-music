package service

import (
	"errors"
	"gf-music/app/api/request"
	"gf-music/app/model/user"
	"github.com/gogf/gf/frame/g"
	uuid "github.com/gogf/guuid"
)

func Register(r *request.Register) (err error) {
	if !user.RecordNotFound(g.Map{"username": r.Username}) {
		return errors.New("用户已存在,注册失败")
	}

	u := user.Entity{
		Id:        0,
		Uuid:      uuid.New().String(),
		Username:  r.Username,
		Password:  r.Password,
		RealName:  "",
		Sex:       0,
		Enable:    0,
		UpdateId:  0,
		CreateId:  0,
		IsAdmin:   0,
		LoginIp:   "",
		Remark:    "",
		Avatar:    "",
		UserEmail: "",
	}
	if err = u.EncryptedPassword(); err != nil { // 哈希加密
		return errors.New("注册失败")
	}
	if _, err = user.Model.Insert(u); err != nil {
		return errors.New("注册失败")
	}
	return
}
