package main

import (
	"fmt"
	"sync"
	"time"
)




/*
    锁 - 资源竞争
1。 按理说最后的结果是0
2. 实际的情况 1. 不是0 2. 每次的运行结果不一样
*/

var wg sync.WaitGroup
var rwlock sync.RWMutex
//1. 互斥锁 2. 读写锁  同步数据
// 能不用锁就不用锁 - 性能
//绝大多数的web系统都是读多写少
//在读和写上加同一把锁
//并发严重下降
//读之间不会产生影响。 写和读之间才会产生影响。 读-写 锁

func read()  {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	rwlock.RUnlock()

}

func write()  {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改成功")
	rwlock.Unlock()

}

func main()  {
	wg.Add(6)
	for i:=0;i<5;i++{
		go read()
	}

	for i:=0;i<1;i++{
		go write()
	}

	wg.Wait()

}

