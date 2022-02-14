package main

import (
	"fmt"
	"time"
)

//goroutine
//python java c++ 多线程 多进程
//go 语言诞生的比较晚, 高并发
//多线程 - 每个线程占用的内存比较多。 系统切换开销比较大 绿程/轻量级线程 - 协程 用户态的线程
//go语言 - 协程
//python两种编程模式： 1. 多线程/多进程  2. 协程并发编程

func p()  {
   for {
   	  fmt.Println("1")
   }
}

//go的协程和python的协程
//python的单个协程占用内存少。
func main()  {
	//fmt.Println("main thread")
	//主死从随
	//go p()
	//闭包
	for i:=0; i<100; i++{
		go func(n int){
			for {
				fmt.Println(i)
				time.Sleep(time.Second * 3)
			}
		}(i)
	}

	time.Sleep(time.Second * 2)

}
