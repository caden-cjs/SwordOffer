/*
@Time :  2020-07-30 15:45
@Author : Caden
@File : main
@Software: GoLand
*/
package main

import (
	"SwordOffer/code/fifo/fifo"
	"fmt"
)

func main() {
	fifo := fifo.NewMyFifoQueueForNum(10)
	for i := 0; i < 10; i++ {
		fifo.Enqueue(i)
	}
	for i := 0; i < 10; i++ {
		dequeue := fifo.Dequeue()
		fmt.Println(dequeue)
	}
}
