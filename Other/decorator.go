package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

/*
@Time    : 2021/3/31 10:25
@Author  : austsxk
@Email   : austsxk@163.com
@File    : decorator.go
@Software: GoLand
*/

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func MyFunc() string {
	return ""
}
func main() {

	//sum1 := timedSumFunc(Sum1)
	//sum2 := timedSumFunc(Sum2)
	//
	//fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
	fmt.Println(getFunctionName(MyFunc))
}
