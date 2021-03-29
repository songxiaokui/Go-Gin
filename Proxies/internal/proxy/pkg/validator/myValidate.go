package validator

import (
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

/*
@Time    : 2021/3/29 15:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : myValidate.go
@Software: GoLand
*/

// 针对自己的模块去注册
func init() {
	if err := myvalidate.RegisterValidation("ProxyUrlValidate", ProxyUrlValidate); err != nil {
		log.Fatal(err)
	}
}

// 自定义需要处理的函数
func ProxyUrlValidate(f1 validator.FieldLevel) bool {
	url, ok := f1.Field().Interface().(string)
	if ok {
		b, err := regexp.Match("^\\w{3,8}$", []byte(url))
		if err != nil {
			return false
		}
		if b {
			return true
		} else {
			return false
		}
	}
	return false
}
