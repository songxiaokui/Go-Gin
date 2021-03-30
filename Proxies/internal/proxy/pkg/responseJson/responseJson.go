package responseJson

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

/*
@Time    : 2021/3/30 09:51
@Author  : austsxk
@Email   : austsxk@163.com
@File    : responseJson.go
@Software: GoLand
*/

// 对gin返回响应进行统一封装,不用写一系列的c.Json(xxx, gin.H{xxx,xxx})等
// 自定义一个jsonResponse

type JSONResponse struct {
	// 自定义的状态码
	Code   string      `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

// 定义一个构造函数
func NewJsonResponse(code string, msg string, result interface{}) *JSONResponse {
	return &JSONResponse{code, msg, result}
}

// 使用Poll,零时对象池，复用该对象，避免重复的创建资源
var ResultPoll *sync.Pool

// 初始化临时对象池
func init() {
	ResultPoll = &sync.Pool{
		New: func() interface{} {
			return NewJsonResponse("", "", "")
		},
	}
}

// 使用闭包修改内部数据
type ResultFunc func(code, msg string, result interface{})

func ResponseOK(c *gin.Context) ResultFunc {
	return func(code, msg string, result interface{}) {
		// 获取对象池中的对象并进行断言
		d := ResultPoll.Get().(*JSONResponse)
		defer ResultPoll.Put(d)
		// 修改其中的值
		d.Code = code
		d.Msg = msg
		d.Result = result
		c.JSON(http.StatusOK, d)
	}
}
