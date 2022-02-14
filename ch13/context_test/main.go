package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1. 监控全局变量
//var stop bool


func cpuInfo()  {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("cpu信息读取完成")

}
func main()  {
	//现在启动一个goroutine去监控某台服务器的CPU的使用情况
	wg.Add(1)
	go cpuInfo()
	//现在希望中断cpu的信息读取

	wg.Wait()
	fmt.Println("信息监控完成")

}
