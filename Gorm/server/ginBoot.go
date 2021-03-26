package server

import (
	"Go-Gin/Gorm/configs"
	"github.com/gin-gonic/gin"
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
	go func() {
		err := s.Run("127.0.0.1:8888")
		if err != nil {
			configs.ErrorChan <- err
		}
	}()
}
