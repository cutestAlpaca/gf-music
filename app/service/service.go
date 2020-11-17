package service

import (
	"errors"
	"gf-music/app/api/request"
	"gf-music/app/model/user"
	"gf-music/library/utils"
	"github.com/gogf/gf/frame/g"
)

var Version = "1.0.0"

func Login(l *request.Login) (data *user.Entity, err error) {
	user := (*user.Entity)(nil) // 用法解释 https://goframe.org/database/gdb/chaining/select#tip4
	userDb := g.DB("default").Table("user").Safe()
	//authorityDb := g.DB("default").Table("authorities").Safe()
	if err = userDb.Where(g.Map{"username": l.Username}).Scan(&user); err != nil {
		return user, errors.New("用户不存在")
	}
	//err = authorityDb.Where(g.Map{"authority_id": user.AuthorityId}).Struct(&user.Authority)
	if utils.CompareHashAndPassword(user.Password, l.Password) { // 检查密码是否正确
		//err = authorityDb.Where(g.Map{"authority_id": user.AuthorityId}).Scan(&user.Authority)
		return user, err
	}
	return user, errors.New("密码错误")
}
