package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
				context.JSON(http.StatusBadRequest, gin.H{"msg": err})
			}
		}()
		context.Next()
	}
}
