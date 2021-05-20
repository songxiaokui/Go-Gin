package main

/*
@Time    : 2021/5/16 23:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

/*
#include <stdio.h>
#include "hello.h"
// 编写自定义C函数与方法
static void SayHello(const char* s) {
	puts(s);
}

// 声明从文件中需要加载的函数
void Say(const char* s);

// 从模块文件中加载
void SayHello2(const char* s);
*/
import "C" // 启用CGO特性

func main() {
	// println("hello cgo")

	// 将Go的字符串转化为C语言的字符串类型
	// C.puts(C.CString("Hello Cgo\n"))

	// C.SayHello(C.CString("hello cgo"))

	// 从文件中调用C定义的函数
	// 不能单独运行，无法加载c文件 需要 go build 或者 go run .
	// C.Say(C.CString("hello RPC"))

	// 从模块文件中加载
	C.SayHello2(C.CString("hello cgo"))
}
