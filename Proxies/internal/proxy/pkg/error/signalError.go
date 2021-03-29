package myerror

/*
@Time    : 2021/3/30 00:07
@Author  : austsxk
@Email   : austsxk@163.com
@File    : signalError.go
@Software: GoLand
*/

// 处理单返回error的姿势
type SignalErrorResult struct {
	err error
}

func (se *SignalErrorResult) Unwrap() {
	if se.err != nil {
		panic(se.err.Error())
	}
}

func MakeSignalError(err error) *SignalErrorResult {
	return &SignalErrorResult{err: err}
}
