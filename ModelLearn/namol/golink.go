package namol

import (
	_ "ModelLearn/runtime"
	_ "unsafe"
)

/*
@Time    : 2021/4/17 14:36
@Author  : austsxk
@Email   : austsxk@163.com
@File    : golink.go
@Software: GoLand
*/

//go:linkname SayHello ModelLearn/runtime.Say
func SayHello(string) string
