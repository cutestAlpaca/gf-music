package service

import (
	"errors"
	"gf-music/app/api/request"
	"gf-music/app/model/user"
	"github.com/ShiningRush/avatarbuilder"
	"github.com/ShiningRush/avatarbuilder/calc"
	"github.com/gogf/gf/frame/g"
	uuid "github.com/gogf/guuid"
	"image/color"
	"math/rand"
	"os"
)

func Register(r *request.Register, ip string) (err error) {
	//if !user.RecordNotFound(g.Map{"username": r.Username}) {
	//	return errors.New("用户已存在,注册失败")
	//}

	u := &user.Entity{
		Id:       0,
		Uuid:     uuid.New().String(),
		Username: r.Username,
		Password: r.Password,
		RealName: "",
		Sex:      r.Sex,
		Enable:   0,
		UpdateId: 0,
		CreateId: 0,
		IsAdmin:  0,
		LoginIp:  ip,
		Remark:   r.Remark,
		Avatar:   "",
		UserEmail: r.Email,
	}
	url := avatar(r.Username);
	path  := "avatar/" + u.Uuid + "/" + u.Username + ".png"
	uploadAvatar, err := uploadFile(url, path); if err != nil {
		g.Log().Println("上传失败")
		return errors.New("上传失败")
	}
	u.ChangeAvatar(uploadAvatar)
	osErr := os.Remove(url)

	if osErr != nil {
		g.Log().Println("删除失败")
		g.Log().Println(osErr)
	}

	if err = u.EncryptedPassword(); err != nil { // 哈希加密
		return errors.New("注册失败")
	}
<<<<<<< HEAD
=======

>>>>>>> 9fe917b75b6123615d240c665a2d6bcb3a47d694
	if _, err = user.Model.Insert(u); err != nil {
		return errors.New("注册失败")
	}
	return
}

// 根据用户名生成头像
func avatar(username string) string {
	imagePath := "./public/resource/image/"
	var colors = []uint32{
		0xff6200, 0x42c58e, 0x5a8de1, 0x785fe0,
	}
	randNum := rand.Intn(4)
	g.Log().Println(randNum)
	ab := avatarbuilder.NewAvatarBuilder("./public/resource/font/SourceHanSansSC-Medium.ttf", &calc.SourceHansSansSCMedium{})
	ab.SetBackgroundColorHex(colors[randNum])
	ab.SetFrontgroundColor(color.White)
	ab.SetFontSize(30)
	ab.SetAvatarSize(200, 200)
	if err := ab.GenerateImageAndSave(username, imagePath + username + ".png"); err != nil {
		g.Log().Println(err)
		g.Log().Panic(err)
	}

	return imagePath + username + ".png"
}
