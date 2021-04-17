package main

import (
	"Go-Gin/Proxies/internal/proxy/pkg/middleware"
	// 此处是一个小细节，验证器初始化的加载在单独的包，所以要主动的引入
	_ "Go-Gin/Proxies/internal/proxy/pkg/validator"
	"Go-Gin/Proxies/internal/proxy/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
@Time    : 2021/3/28 21:22
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

func main() {
	s := gin.Default()
	// 使用包装简化路由等
	server.NewMyEngine(s).SayHelloHandler()

	v1 := s.Group("/v1/proxies")
	{
		v1.Use(middleware.ErrorMiddleware())
		// 获取资源列表,已经添加了v1组的前缀了,可以添加query参数
		//v1.GET("", server.ProxiesListHandlerFunc())
		v1.GET("", server.ProxiesListHandler)

		// 根据id获取资源详情
		v1.GET("/:id", server.GetProxyByIdHandler)

		// 添加权限认证
		v1.Use(middleware.AuthenticationMiddleware())
		{
			// 添加一个资源
			v1.POST("", server.CreateProxyHandler)
			// 删除一个资源
			v1.DELETE("/:id", server.DeleteProxyHandler)
		}

	}
	s.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("hello gin"))
	})
	errChan := make(chan error)
	signalChan := make(chan os.Signal)
	go func() {
		if err := s.Run(":8888"); err != nil {
			errChan <- err
		}
	}()
	go func() {
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)
		errChan <- fmt.Errorf("%s", <-signalChan)
	}()

	select {
	case err := <-errChan:
		log.Fatal(err)
	}
}
