package server

import "Go-Gin/Gorm/configs"

/*
@Time    : 2021/3/26 21:03
@Author  : austsxk
@Email   : austsxk@163.com
@File    : dbBoot.go
@Software: GoLand
*/

func DbBoot() {
	go func() {
		err := configs.InitDb()
		if err != nil {
			configs.ErrorChan <- err
		}
	}()
}
