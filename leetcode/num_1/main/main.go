/*
@Time : 2020/04/25 21:21
@Author : caden
@File : main.go
@Software: GoLand
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Printf("twoSumOne=%v\n", twoSumOne(ints, target))
	fmt.Printf("twoSumTwo=%v\n", twoSumTwo(ints, target))
}
func twoSumOne(nums []int, target int) []int {
	for index, item := range nums {
		num2 := target - item
		for i := index + 1; i < len(nums); i++ {
			if nums[i] == num2 {
				return []int{index, i}
			}
		}
	}
	return nil
}
func twoSumTwo(nums []int, target int) []int {
	myMap := make(map[int]int)
	for index, item := range nums {
		value, exist := myMap[target-item]
		if exist {
			return []int{value, index}
		}
		myMap[item] = index
	}
	return nil
}
