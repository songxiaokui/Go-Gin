package main

/*
@Time    : 2021/3/26 20:56
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	gc "Go-Gin/Gorm/configs"
	gs "Go-Gin/Gorm/server"
	"log"
)

func main() {
	gs.DbBoot()
	gs.GinBoot()
	err := <-gc.ErrorChan
	if err != nil {
		log.Fatal(err)
	}
}
