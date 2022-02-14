package main

import (
	"fmt"
	"time"
)

func main()  {
	/*
	  go 语言提供了select 功能 作用于channel之上的。 多路复用
	select 会随机公平的选择一个case语句执行
	select的应用场景 1. timeout的超时机制
	*/
	//timeout := false

	timeout := make(chan bool, 1)
	go func() {
		//该 goroutine如果执行时间超过1s， 那么就报告主的goroutine
		time.Sleep(time.Second*2)
		timeout <- true
	}()

	select {
	case <- timeout:
		fmt.Println("超时了")

	}
	//fmt.Println(<- timeout)

	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 2)
	//
	//ch1 <- 1
	//ch2 <- 2
	//
	//select {
	//case data := <- ch1:
	//	fmt.Println(data)
	//case data := <- ch2:
	//	fmt.Println(data)
	//}

}
