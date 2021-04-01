package middleware

import (
	myrsp "Go-Gin/Proxies/internal/proxy/pkg/responseJson"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
@Time    : 2021/3/29 23:16
@Author  : austsxk
@Email   : austsxk@163.com
@File    : error.go
@Software: GoLand
*/

// 处理error的middleware
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 捕获异常
		defer func() {
			if err := recover(); err != nil {
				// 修改使用封装好的包
				msg := "请求错误"
				if m, ok := err.(string); ok {
					msg = m
				}
				myrsp.R(context)(strconv.Itoa(http.StatusBadRequest), msg, nil)(myrsp.Failed)
			}
		}()
		context.Next()
	}
}
