package service

import (
	"context"
	"errors"
	"fmt"
	"gf-music/app/api/request"
	"gf-music/app/model/user"
	"github.com/ShiningRush/avatarbuilder"
	"github.com/ShiningRush/avatarbuilder/calc"
	"github.com/gogf/gf/frame/g"
	uuid "github.com/gogf/guuid"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"image/color"
	"math/rand"
)

func Register(r *request.Register, ip string) (err error) {
	//if !user.RecordNotFound(g.Map{"username": r.Username}) {
	//	return errors.New("用户已存在,注册失败")
	//}

	u := user.Entity{
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
		Remark:   "",
		//Avatar:    uploadFile("/home/koko/Downloads/docker.png"),
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

// 不行就前端生成随即头像
func avatar(username string) string {
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
	if err := ab.GenerateImageAndSave(username, "./public/resource/image/ "+username+".png"); err != nil {
		fmt.Println(err)
		return "err"
	}
	//defer uploadFile("./public/resource/image/ "+username+".png", username)
	return "aa"
}

// 上传七牛后删除
func uploadFile(url string) string {
	var (
		accessKey = "dp50sTuARMeOk6FBC9GruqGtdAdO0X3-jMR3wHqj" // 七牛的accessKey 去七牛后台获取
		secretKey = "gx0kITZdk-0tIPTsJyza3Q1iNXiFzJadHrCFR9i8" // 七牛的secretKey 去七牛后台获取
		bucket    = "gf-music"                                 // 上传空间 去七牛后台创建
	)
	// 鉴权
	mac := qbox.NewMac(accessKey, secretKey)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuabei //指定上传的区域
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      //是否使用CDN上传加速

	// 需要上传的文件
	localFile := url

	// 七牛key
	qiniuKey := "qiniu/20200113/test.png"

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, localFile, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		return "err"
	}
	fmt.Println("上传成功,key为:", ret.Key)
	return "err2"
}
