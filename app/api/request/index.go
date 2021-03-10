package request

// Login request Structure
type Login struct {
	Username string `p:"username" v:"required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
	//Captcha   string `json:"captcha" valid:"required#请输入正确的验证码"`
	CaptchaId string `json:"captchaId" valid:"required|length:5,20#请输入captchaId|您输入captchaId长度非法"`
}

// Register request Structure
type Register struct {
	Username string `p:"username" v:"required|length:1,30#请输入用户名称|账号长度为:min到:max位"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
	Sex      int    `p:"sex"`
	Remark string `p:"remark"`
	Email   string `p:"email" v:"email|length:1,30#请输入邮箱|头像地址长度为:min到:max位"`
	//HeaderImg   string `p:"headerImg"`
	//AuthorityId string `p:"authorityId" v:"required|length:1,100#请输入密码|authority_id长度为:min到:max位"`
}
