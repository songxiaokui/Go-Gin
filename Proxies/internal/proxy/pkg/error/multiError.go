package myerror

import "errors"

/*
@Time    : 2021/3/30 00:19
@Author  : austsxk
@Email   : austsxk@163.com
@File    : multiError.go
@Software: GoLand
*/
// 其实包含了一种参数的验证方法
type MultiErrorResult struct {
	data interface{}
	err  error
}

func (me *MultiErrorResult) Unwrap() interface{} {
	if me.err != nil {
		panic(me.err.Error())
	}
	return me.data
}

// 可变参数，可能是一个返回值，可能是两个返回值
func MakeMultiError(vs ...interface{}) *MultiErrorResult {
	if len(vs) == 1 {
		if vs[0] == nil {
			return &MultiErrorResult{nil, nil}
		}
		if d, ok := vs[0].(error); ok {
			return &MultiErrorResult{nil, d}
		}
	}
	if len(vs) == 2 {
		if vs[1] == nil {
			return &MultiErrorResult{vs[0], nil}
		}
		if d, ok := vs[1].(error); ok {
			return &MultiErrorResult{vs[0], d}
		}
	}
	return &MultiErrorResult{nil, errors.New("params format error")}
}
