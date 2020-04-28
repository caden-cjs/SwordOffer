/*
@Time : 2020/04/25 21:50
@Author : caden
@File : main.go
@Software: GoLand
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 8
	fmt.Printf("1+2+.....+%v=%v\n", n, sumNumsOne(n))
	fmt.Printf("1+2+.....+%v=%v\n", n, sumNumsThree(n))
}
func sumNums(n int) int {
	var f func(ret *int, n int) bool
	f = func(ret *int, n int) bool {
		*ret += n
		return n != 0 && f(ret, n-1)
	}

	var ans int
	f(&ans, n)

	return ans
}

func sumNumsOne(n int) int {
	var ans int
	var fu func(rest *int, n int) bool
	fu = func(rest *int, n int) bool {
		*rest += n
		return n != 0 && fu(rest, n-1)
	}
	fu(&ans, n)
	return ans
}
func sumNumsThree(n int) int {
	slot := make(chan struct{}, n)
	defer close(slot)
	sum := 0
loop:
	sum += n
	n -= 1
	select {
	case slot <- struct{}{}:
	default:
		return sum
	}
	goto loop
}
func sumNumsFour(n int) int {
	return (int(math.Pow(float64(n), float64(2))) + n) >> 1
}
