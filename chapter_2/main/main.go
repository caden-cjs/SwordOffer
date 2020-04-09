/*
@Time : 2020/04/08 22:23
@Author : caden
@File : main.go
@Software: GoLand
*/
package main

func main() {
	var arrs [][]int = [][]int{
		{1, 2, 3, 4},
		{3, 4, 5, 8},
		{5, 7, 8, 11},
		{7, 9, 12, 14},
	}
}
func querySevenOrFire(arrs [][]int) bool {
	width := len(arrs)
	len := len(arrs[0])

	indexX, indexY := (width-1)/2, (len-1)/2 //寻找一个中间点
	i := arrs[indexX][indexY]
	if i >= 5 {
		for x, y := indexX, indexY; x == 0 || y == 0; {

		}
	} else {

	}
}
