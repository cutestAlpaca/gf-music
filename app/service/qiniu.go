package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 上传七牛后删除
func uploadFile(url, filePath string) (string, error) {
	var (
		accessKey = g.Cfg("private").GetString("qiniu.AccessKey") // 七牛的accessKey 去七牛后台获取
		secretKey = g.Cfg("private").GetString("qiniu.SecretKey") // 七牛的secretKey 去七牛后台获取
		bucket    = g.Cfg("private").GetString("qiniu.Bucket")                              // 上传空间 去七牛后台创建
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
	cfg.UseHTTPS = g.Cfg("private").GetBool("qiniu.UseHTTPS")           // 是否使用https域名
	cfg.UseCdnDomains = g.Cfg("private").GetBool("qiniu.UseCdnDomains")      //是否使用CDN上传加速

	// 需要上传的文件
	localFile := url

	// 七牛key
	qiniuKey := filePath

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 上传文件
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, localFile, nil)
	if err != nil {
		g.Log().Println("上传文件失败,原因:", err)
		return "err", err
	}

	g.Log().Println("上传成功,key为:", ret)
	return ret.Key, nil
}

func AddUrl(url string) string {
	return g.Cfg("private").GetString("qiniu.Url") + url
}