package utils

import (
	"fmt"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"net"
	"time"
)

const PublicKey = "HqmP1KLMuz09Q0Bu"

//服务端ip
func GetLocalIP() (ip string, err error) {
	address, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range address {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

//获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

//获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

//日期字符串转时间戳（秒）
func StrToTimestamp(dateStr string) int64 {
	tm, err := gtime.StrToTime(dateStr)
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return tm.Timestamp()
}

//时间戳转 yyyy-MM-dd HH:mm:ss
func TimeStampToDateTime(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d H:i:s")
}

//时间戳转 yyyy-MM-dd
func TimeStampToDate(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d")
}

//获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := ghttp.GetBytes(url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.GetInt("code") == 0 {
		city := json.GetString("city")
		return city
	} else {
		return ""
	}
}

//获取当前请求接口域名
func GetDomain(r *ghttp.Request) (string, error) {
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"]), nil
}

//字符串加密
func EncryptCBC(plainText, publicKey string) string {
	key := []byte(publicKey)
	b, e := gaes.EncryptCBC([]byte(plainText), key, key)
	if e != nil {
		g.Log().Error(e.Error())
		return ""
	}
	return gbase64.EncodeToString(b)
}

// 微信通知
func WeChatNotification(title, content string) bool {
	if res, err := g.Client().Post(g.Cfg("private").GetString("wechat.url"), g.Map{
		"title":   title,
		"content": content,
	}); err != nil {
		g.Log().Printf("%v\n", res)
		return false
	} else {
		if res.Response.StatusCode == 200 {
			return true
		}
		defer res.Close()
	}

	return true
}
