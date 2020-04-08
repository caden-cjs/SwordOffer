/*
@Time : 2020/04/08 21:58
@Author : caden
@File : main.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < 10; i++ {
		/*1*/
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			fmt.Println(m)
			mu.Unlock()
		}(i)
		/*2*/
		//go func() {
		//	defer wg.Done()
		//	mu.Lock()
		//	m[i] = i
		//	fmt.Println(m)
		//	mu.Unlock()
		//}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
