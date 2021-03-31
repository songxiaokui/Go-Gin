package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

/*
@Time    : 2021/3/31 10:30
@Author  : austsxk
@Email   : austsxk@163.com
@File    : httpDecorator.go
@Software: GoLand
*/

// http 各种认证装饰器
// With 方式装饰,装饰谁,返回谁

// 定义一个需要处理业务的路由
func HellO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, decorator!"+r.URL.Path)
}

// 添加各种装饰
// 使用闭包实现,装饰谁返回谁
func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

func WithBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithBasicAuth()")
		cookie, err := r.Cookie("Auth")
		if err != nil || cookie.Value != "Pass" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func WithDebugLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithDebugLog")
		r.ParseForm()
		log.Println(r.Form)
		log.Println("path", r.URL.Path)
		log.Println("scheme", r.URL.Scheme)
		log.Println(r.Form["url_long"])
		for k, v := range r.Form {
			log.Println("key:", k)
			log.Println("val:", strings.Join(v, ""))
		}
		h(w, r)
	}
}

// pipeline方式改造, 函数签名就是改造的方式
type DecoratorFunc func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Handler(fn http.HandlerFunc, opt ...DecoratorFunc) http.HandlerFunc {
	// 遍历装饰的函数
	for _, decorator := range opt {
		fn = decorator(fn)
	}
	return fn
}

// 上面的装饰器不具备泛型,使用interface{} + reflection可以实现泛型
// 装饰前需要提前声明一个类型

func Decorators(decoPtr, fn interface{}) (err error) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return
}

func Decorator(decoratorPtr, fn interface{}) (err error) {
	// 参数是装饰的函数类型 装饰的函数
	var decoratedFunc, targetFunc reflect.Value
	// interface类型，获取要用Elem()
	decoratedFunc = reflect.ValueOf(decoratorPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	// 创造一个函数，其类型为目标类型
	newFunc := reflect.MakeFunc(decoratedFunc.Type(), func(args []reflect.Value) (results []reflect.Value) {
		fmt.Println("before logic")
		// call real logic, target method
		results = targetFunc.Call(args)
		// after logic
		fmt.Println("after logic")
		return
	})
	decoratedFunc.Set(newFunc)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b, c)
	return a + b + c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return a + b
}

func main() {
	// 使用多重装饰即可，但是不好看
	// http.HandleFunc("/", WithDebugLog(WithDebugLog(HellO)))

	// 使用pipeline方式实现
	// D := []DecoratorFunc{WithDebugLog, WithBasicAuth, WithAuthCookie}
	// http.HandleFunc("/hello", Handler(HellO, D...))

	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	err := Decorator(&myfoo, foo)
	myfoo(1, 2, 3)
	fmt.Println(err)

	// 先声明一个装饰的函数类型
	type MyBar func(string, string) string
	var mybar MyBar
	err = Decorator(&mybar, bar)
	mybar("hello", "golang")
	fmt.Println(err)

	/*
		Decorator 这个函数其实是可以修饰几乎所有的函数的。于是，这种可以通用于其它函数的编程方式，
		可以很容易地将一些非业务功能的、属于控制类型的代码给抽象出来（所谓的控制类型的代码就是像
		for-loop，或是打日志，或是函数路由，或是求函数运行时间之类的非业务功能性的代码）
	*/
}
