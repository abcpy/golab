
package main

import (
"fmt"
"sync"
"time"
)

var wg sync.WaitGroup

//deadlock

func comsumer(queue chan int)  {
	defer wg.Done()
	//data := <- queue
	//data1 := <- queue
	//for data:= range queue{
	//	fmt.Println(data)
	//	time.Sleep(time.Second)
	//}

	for {
		data, ok := <- queue
		if !ok {
			break
		}
		fmt.Println(data, ok)
		time.Sleep(time.Second)
		queue <- 2
	}

}

func main()  {
	/*
	   channel 类型 1. 有缓冲 2. 没有缓冲
	    双向的 单向的
	*/
	//定义一个channel
	//var msg chan int
	var msg chan <- int   //单向的channel
	//var msg <- chan int
	//初始化channel 两种方式 1. 无缓冲 2. 有缓冲
	//msg = make(chan int)
	//msg = make(chan int, 1) // 有缓冲空间的
	msg = make(chan int, 10) // 有缓冲空间的
	//go语言中使用make初始化的有三种 1。 slice 2. map 3. channel

	//msg <- 1 //将1 放入channel中
	msg <- 2
	wg.Add(1)
	//data := <- msg
	//go comsumer(msg)
	msg <- 1

	//msg <- 2
	//关闭channel 已经关闭的channel不能在发送数据了  已经关闭的channel 能够继续取数据
	fmt.Println("等待返回值")
	//fmt.Println(<-msg)
	close(msg)
	wg.Wait()
	//msg <- 3 //panic: send on closed channel


	//msg <- 2 //将1 放入channel中


	//data := <- msg
	//
	//fmt.Println(data)



}

