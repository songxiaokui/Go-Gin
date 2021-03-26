package configs

/*
@Time    : 2021/3/26 20:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : CommonConfig.go
@Software: GoLand
*/
// 定义一个错误的channel 异步处理错误
var ErrorChan chan error

func init() {
	ErrorChan = make(chan error)
}
