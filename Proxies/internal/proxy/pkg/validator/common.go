package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

/*
@Time    : 2021/3/29 17:29
@Author  : austsxk
@Email   : austsxk@163.com
@File    : common.go
@Software: GoLand
*/

var myvalidate *validator.Validate

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		myvalidate = v
	} else {
		log.Fatal("validate binding error")
	}
}
