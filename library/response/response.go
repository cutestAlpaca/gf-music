package response

import "github.com/gogf/gf/net/ghttp"

type JsonResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode int = 0
	ErrorCode   int = -1
)

// 标准返回结果数据结构封装
func Json(r *ghttp.Request,code int,message string, data ...interface{})  {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(JsonResponse{
		Code: code,
		Message: message,
		Data: responseData,
	})
}

// 返回正确状态
func JsonSuccess(r *ghttp.Request, message string, data ...interface{})  {
	r.Response.WriteJson(JsonResponse{
		Code:    200,
		Message: message,
		Data:    data,
	})
	r.Exit()
}

// 返回JSON数据并退出当前HTTP执行函数
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}