/*
@Time : 2020/04/28 19:49
@Author : caden
@File : mian_test.go.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"math"
	"testing"
)

func Test_coinChange(t *testing.T) {
	ints := []int{1, 2, 5}
	amount := 11
	change := coinChange(ints, amount)
	fmt.Println(change)
}

/**
凑硬币
*/
func coinChange(coins []int, amount int) int {
	result := math.MaxInt64
	myMap := make(map[int]int, 0)
	var fun func(coins []int, amount int, num int)
	fun = func(coins []int, amount int, num int) {
		_, ok := myMap[amount]
		if !ok {
			myMap[amount] = num
		} else {
			return
		}
		for _, item := range coins {
			if item == amount {
				fmt.Printf("遍历次数=%v,result=%v\n", num+1, result)
				result = int(math.Min(float64(num+1), float64(result)))
			}
			am := amount - item
			if am > 0 {
				fun(coins, am, num+1)
			}
		}
	}
	fun(coins, amount, 0)
	fmt.Println(myMap)
	return result
}
