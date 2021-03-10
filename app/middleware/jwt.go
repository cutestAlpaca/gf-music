package middleware

import (
	"fmt"
	v1 "gf-music/app/api/v1"
	"gf-music/library/global"

	jwt "github.com/gogf/gf-jwt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// JwtAuth 验证token有效性
func JwtAuth(r *ghttp.Request) {
	v1.GfJWTMiddleware.MiddlewareFunc()(r)
	Token, err := v1.GfJWTMiddleware.ParseToken(r) // 解析token
	fmt.Printf("token :%v\n", Token)
	if err != nil {
		if err == jwt.ErrExpiredToken {
			global.Result(r, global.ERROR, g.Map{"reload": true}, "授权已过期")
			r.ExitAll()
		}
		global.Result(r, global.ERROR, g.Map{"reload": true}, err.Error())
		r.ExitAll()
	}
	//token := Token.Raw
	//if service.IsBlacklist(token) {
	//	global.Result(r, global.ERROR, g.Map{"reload": true}, "您的帐户异地登陆或令牌失效")
	//	r.ExitAll()
	//}
	var claims = gconv.Map(Token.Claims)
	r.SetParam("claims", Token.Claims)
	r.SetParam("user_authority_id", claims["user_authority_id"])
	//if g.Cfg().GetBool("system.UseMultipoint") {
	//	if !ValidatorRedisToken(gconv.String(claims["admin_uuid"]), token) {
	//		global.FailWithMessage(r, "Token鉴权失败")
	//		r.Exit()
	//	}
	//}
	r.Middleware.Next()
}

// GetToken 根据uuid获取
// 抑制了错误, 但是不建议使用
func GetToken(userUUID string) string {
	conn := g.Redis().Conn()
	defer conn.Close()
	r, _ := conn.Do("GET", userUUID)
	return gconv.String(r)
}

// ValidatorRedisToken 鉴权jwt
func ValidatorRedisToken(userUUID string, oldToken string) bool {
	if GetToken(userUUID) != oldToken {
		return false
	}
	return true
}
