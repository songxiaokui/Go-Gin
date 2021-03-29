package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
@Time    : 2021/3/29 09:51
@Author  : austsxk
@Email   : austsxk@163.com
@File    : middleware.go
@Software: GoLand
*/

// 鉴权中间键、确保用户是登陆状态
// 使用，给需要认证的路由，添加Use，将该中间键传入 s.Use(AuthenticationMiddleware())
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 此处进行jwt验证或则token验证
		// 暂时使用token获取认证
		if token, ok := c.GetQuery("token"); ok {
			if token == "token" {
				c.Next()
				return
			} else {
				goto shutDown
			}

		}
	shutDown:
		// 只要认证没通过,抛出异常
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Unauthorized"})
		c.Abort()
	}
}
