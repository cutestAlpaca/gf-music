package service

import (
	"database/sql"
	"errors"
	"gf-music/app/model/sys_login_log"
	"gf-music/app/model/user"
	"gf-music/app/model/user_online"
	"gf-music/library/response"
	"gf-music/library/utils"
	"github.com/goflyfox/gtoken/gtoken"
	_ "github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/user_agent"
	"strings"
)

const Version  = "1.0.0"

//获取字母数字混合验证码
func GetVerifyImgString() (idKeyC string, base64stringC string) {
	driver := &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
	}
	driver = driver.ConvertFonts()
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	idKeyC, base64stringC, err := c.Generate()
	if err != nil {
		g.Log().Error(err)
	}
	return
}

//验证输入的验证码是否正确
func VerifyString(id, answer string) bool {
	driver := new(base64Captcha.DriverString)
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}

// 登录返回方法
func LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		token := respData.GetString("token")
		uuid := respData.GetString("uuid")
		var userInfo *user.Entity
		r.GetParamVar("userInfo").Struct(&userInfo)
		//保存用户在线状态token到数据库
		userAgent := r.Header.Get("User-Agent")
		ua := user_agent.New(userAgent)
		os := ua.OS()
		explorer, _ := ua.Browser()
		entity := user_online.Entity{
			Uuid:       uuid,
			Token:      token,
			CreateTime: gconv.Uint64(gtime.Timestamp()),
			UserName:   userInfo.UserName,
			Ip:         utils.GetClientIp(r),
			Explorer:   explorer,
			Os:         os,
		}
		entity.Save()
		r.Response.WriteJson(gtoken.Succ(g.Map{
			"token": token,
		}))
	}
}

//gtoken验证后返回
func AuthAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if r.Method == "OPTIONS" || respData.Success() {
		r.Middleware.Next()
	} else {
		respData.Msg = "用户信息验证失败"
		response := r.Response
		options := response.DefaultCORSOptions()
		response.CORS(options)
		response.WriteJson(respData)
		r.ExitAll()
	}
}


//后台退出登陆
func LoginOut(r *ghttp.Request) bool {
	//删除在线用户状态
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" && parts[1] != "" {
			//删除在线用户状态操作
			user_online.Model.Delete("token", parts[1])
		}
	}
	authHeader = r.GetString("token")
	if authHeader != "" {
		//删除在线用户状态操作
		user_online.Model.Delete("token", authHeader)
	}
	return true
}

// 用户登录，成功返回用户信息，否则返回nil
func signIn(username, password string, r *ghttp.Request) (error, *user.Entity) {
	user, err := user.Model.Where("user_name=? and user_password=?", username, password).One()
	if err != nil && err != sql.ErrNoRows {
		return err, nil
	}
	if user == nil {
		return errors.New("账号或密码错误"), nil
	}
	//判断用户状态
	if user.UserStatus == 0 {
		return errors.New("用户已被冻结"), nil
	}
	returnData := *user
	//更新登陆时间及ip
	user.LastLoginTime = gconv.Int(gtime.Timestamp())
	user.LastLoginIp = utils.GetClientIp(r)
	user.Update()
	return nil, &returnData
}

//登录日志记录
func loginLog(status int, username, ip, userAgent, msg, module string) {
	var log sys_login_log.Entity
	log.LoginName = username
	log.Ipaddr = ip
	log.LoginLocation = utils.GetCityByIp(log.Ipaddr)
	ua := user_agent.New(userAgent)
	log.Browser, _ = ua.Browser()
	log.Os = ua.OS()
	log.Status = status
	log.Msg = msg
	log.LoginTime = gtime.Timestamp()
	log.Module = module
	log.Save()
}


//AdminLogin 后台用户登陆验证
func AdminLogin(r *ghttp.Request) (string, interface{}) {

	data := r.GetFormMapStrStr()
	rules := map[string]string{
		"idValueC": "required",
		"username": "required",
		"password": "required",
	}
	msgs := map[string]interface{}{
		"idValueC": "请输入验证码",
		"username": "账号不能为空",
		"password": "密码不能为空",
	}

	if e := gvalid.CheckMap(data, rules, msgs); e != nil {
		response.JsonExit(r, response.ErrorCode, e.String())
	}
	//判断验证码是否正确
	if !VerifyString(data["idKeyC"], data["idValueC"]) {
		response.JsonExit(r, response.ErrorCode, "验证码输入错误")
	}
	password := utils.EncryptCBC(data["password"], utils.PublicKey)
	var keys string
	keys = data["username"] + password

	ip := utils.GetClientIp(r)
	userAgent := r.Header.Get("User-Agent")
	if err, user := signIn(data["username"], password, r); err != nil {
		go loginLog(0, data["username"], ip, userAgent, err.Error(), "系统后台")
		response.JsonExit(r, response.ErrorCode, err.Error())
	} else {
		//判断是否后台用户
		if user.IsAdmin != 1 {
			response.JsonExit(r, response.ErrorCode, "抱歉!此用户不属于后台管理员!")
		}
		r.SetParam("userInfo", user)
		go loginLog(1, data["username"], ip, userAgent, "登录成功", "系统后台")
		return keys, user
	}
	return keys, nil
}