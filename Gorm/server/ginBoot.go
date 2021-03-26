package server

import (
	"Go-Gin/Gorm/biz"
	"Go-Gin/Gorm/configs"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
@Time    : 2021/3/26 20:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : ginBoot.go
@Software: GoLand
*/

func GinBoot() {
	s := gin.New()
	s.GET("/", func(g *gin.Context) {
		g.JSON(200, gin.H{"message": "ok"})
	})
	s.GET("/:proxy", func(context *gin.Context) {
		id := context.Param("proxy")
		ids, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(404, gin.H{"message": "Not Found"})
		}
		context.JSON(200, biz.NewProxies().LoadById(ids))
	})
	go func() {
		err := s.Run("127.0.0.1:8888")
		if err != nil {
			configs.ErrorChan <- err
		}
	}()
}
