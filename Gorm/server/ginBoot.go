package server

import (
	"Go-Gin/Gorm/biz"
	"Go-Gin/Gorm/configs"
	"fmt"
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

	// 测试where查询,非基类
	s.GET("/proxy/filter", func(context *gin.Context) {
		opt, _ := context.GetQuery("opt")
		value, _ := context.GetQuery("value")
		o, err := strconv.Atoi(opt)
		if err != nil {
			context.JSON(404, gin.H{"message": "Operate type error"})
			return
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			context.JSON(404, gin.H{"message": "Value type error"})
			return
		}
		// 生成过滤条件
		p := biz.NewProxies()
		p.Filter(p.IdCompareGenerate(v, o))
		fmt.Println("数据库查询的结果: ", p)
		context.JSON(200, p)

	})

	// 使用通用基类方法查询
	s.GET("/proxy/filters", func(context *gin.Context) {
		opt, _ := context.GetQuery("opt")
		value, _ := context.GetQuery("value")
		o, err := strconv.Atoi(opt)
		if err != nil {
			context.JSON(404, gin.H{"message": "Operate type error"})
			return
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			context.JSON(404, gin.H{"message": "Value type error"})
			return
		}
		// 生成过滤条件
		p := biz.NewProxies()
		// 由于子类也实现类父类，显示的使用父亲的方法
		p.BaseModel.Filter(p.BaseModel.CompareByFiled("id", v, o))
		fmt.Println("数据库查询的结果: ", p)
		context.JSON(200, p)
	})

	s.GET("/proxies/:id", func(context *gin.Context) {
		id := context.Param("id")
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
