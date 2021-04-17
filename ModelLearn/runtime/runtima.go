package runtime

import (
	"fmt"
	_ "unsafe"
)

/*
@Time    : 2021/4/17 14:36
@Author  : austsxk
@Email   : austsxk@163.com
@File    : runtima.go
@Software: GoLand
*/

func Say(s string) string {
	return fmt.Sprintf("you say: %s", s)
}
