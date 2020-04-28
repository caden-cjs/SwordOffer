/*
@Time : 2020/04/28 19:05
@Author : caden
@File : mian.go
@Software: GoLand
*/
package main

import "fmt"

func main() {
	num := lib(100)
	fmt.Print(num)
}

func lib(n int) int {
	ints := make([]int, 3)
	var f func(intSlice []int, currNum int)
	f = func(intSlice []int, currNum int) {
		if currNum == 1 || currNum == 2 {
			ints[currNum] = 1
		} else {
			intSlice[currNum%3] = intSlice[(currNum-1)%3] + intSlice[(currNum-2)%3]
		}
		if currNum == n {
			return
		} else {
			f(ints, currNum+1)
		}
	}
	f(ints, 1)
	return ints[n%3]
}


