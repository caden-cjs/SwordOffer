/*
@Time : 2020/04/08 21:03
@Author : caden
@File : main.go
@Software: GoLand
*/
package main

import (
	"SwordOffer/chapter_1/module"
	"fmt"
	"sync"
	"time"
)

func main() {
	ints := make(chan int, 10)
	defer close(ints)
	group2 := &sync.WaitGroup{}
	group2.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ints <- i
		}
		group2.Done()
		close(ints)
	}()
	group2.Wait()
	group := &sync.WaitGroup{}
	group.Add(1)
	for i := 1; i <= 20; i++ {
		go func() {
			select {
			case x := <-ints:
				{
					for {
						//x%2==1)?&module.Person{}:&module.Student{}
						var typeT interface{}
						if x%2 == 1 {
							typeT = &module.Person{}
						} else {
							typeT = &module.Student{}
						}
						singleton := module.GetSingleton(typeT)
						fmt.Printf("%v获取到的Person地址=%p\n", x, singleton)
						time.Sleep(time.Millisecond * 1300)
					}
				}
			default:
				{
					fmt.Printf("获取完毕了\n")
					return
				}
			}
		}()
		time.Sleep(time.Millisecond * 1000)
	}
	group.Wait()
	close(ints)
	//singleton := module.GetSingleton(&module.Person{})
	//person := singleton.(*module.Person)
	//singleton2 := module.GetSingleton(&module.Person{})
	//person2 := singleton2.(*module.Person)
	//singleton3 := module.GetSingleton(&module.Person{})
	//person3 := singleton3.(*module.Person)
	//fmt.Printf("person的地址=%p\t,person2的地址=%p\t,person3的地址=%p", person, person2, person3)
}
