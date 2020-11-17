package service

import (
	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"
)

// GetRedisJWT Get user info in redis
// GetRedisJWT 获取用户在Redis的token
func GetRedisJWT(userUUID string) (string, error) {
	conn := g.Redis().Conn()
	defer conn.Close()
	r, err := conn.Do("GET", userUUID)
	return gconv.String(r), err
}

// SetRedisJWT set jwt into the Redis
// SetRedisJWT 保存jwt到Redis
func SetRedisJWT(userUUID string, jwt string) (err error) {
	_, err = g.Redis().Do("SETEX", userUUID, g.Cfg("jwt").GetUint("jwt.ExpiresAt")*3600000000000, jwt)
	return err
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
