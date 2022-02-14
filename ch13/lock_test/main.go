package main

import (
	"fmt"
	"sync"
)

/*
    锁 - 资源竞争
1。 按理说最后的结果是0
2. 实际的情况 1. 不是0 2. 每次的运行结果不一样
*/

var total int
var wg sync.WaitGroup
var lock sync.Mutex
//1. 互斥锁 2. 读写锁  同步数据
// 能不用锁就不用锁 - 性能
//绝大多数的web系统都是读多写少
//在读和写上加同一把锁
//并发严重下降
//读之间不会产生影响。 写和读之间才会产生影响。 读-写 锁

func add()  {
	defer wg.Done()
	for i:=0; i<10000;i++{
		// 先锁上
		lock.Lock()
		total = total + 1  //这个代码和total-1并发运行
		lock.Unlock()
		//
		//1. 从total取出值
		//2. 将total + 1
		//3. 将total + 1的结果赋值给total
	}

}

func sub()  {
	defer wg.Done()
	for i:=0; i<10000;i++{
		lock.Lock()
		total = total - 1
		lock.Unlock()
	}

}
func main()  {
	wg.Add(2)
   go add()
	go sub()
	wg.Wait()
	fmt.Println(total)

}
