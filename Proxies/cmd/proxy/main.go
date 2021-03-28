package main

import (
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
