package main

/*
@Time    : 2021/5/18 09:34
@Author  : austsxk
@Email   : austsxk@163.com
@File    : hello.go
@Software: GoLand
*/

// 使用go实现C的标准接口

import "C"

import (
	"fmt"
)

// 将Go语言实现的函数SayHello导出为C语言函数
// export SayHello2
func SayHello2(s *C.char) {
	fmt.Println(C.GoString(s))
}
