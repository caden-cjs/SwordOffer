/*
@Time :  2020-08-13 16:34
@Author : Caden
@File : main
@Software: GoLand
*/
package main

import "fmt"

func main() {
	fmt.Println(strStr("", ""))
}

func strStr(haystack string, needle string) int {
	if len(haystack)==0 || len(needle)==0{
		return -1
	}
	index := -1
	for i, _ := range needle {
		for j, item := range haystack {
			if item == int32(needle[i]) {
				if index == -1 {
					index = j
				}
				i++
				if i == len(needle) {
					return index
				}
			} else {
				index = -1
				continue
			}

		}
		return -1
	}
	return index
}
func strStr2(haystack string, needle string) int {
	for i := 0; (len(haystack) - i) > len(needle); i++ {
		for j := 0; j < len(needle); i++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if len(needle) == j {
				return i
			}
		}
	}
	return -1
}
