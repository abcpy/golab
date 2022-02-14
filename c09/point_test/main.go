package main

import "fmt"

func main()  {
	//指针 -- 对于内存来说， 每一个字节都有地址-- 通过16进制打印出来
	a := 10
	fmt.Printf("%p\n", &a) //0xc00000a0b0 变量有地址
	//这个变量只能保存地址值
	var ip *int //这个变量只能保存地址类型这种值
	ip = &a
	fmt.Println(ip) //0xc00000a0b0

	//如何定义指针变量， 修改指针变量指向的内存中的值
	*ip = 30
	fmt.Println(a) //30
	fmt.Println(*ip) //30
	fmt.Printf("变量所在内存的地址：%p, 变量的值: %v\n", ip, *ip) //变量所在内存的地址：0xc00000a0b0, 变量的值: 30
     //在python中list和dict都是引用传递

     //指针可以指向数组，指向数组的指针 数组是值类型
      arr := [3]int{1, 2, 3}

      var ip2 *[3]int = &arr
      fmt.Println(ip2)

      var ptrs [3]*int // 创建能够存储三个指针变量的数组
      fmt.Println(ptrs) //[<nil> <nil> <nil>]

      //指针的默认值是nil

      //make new nil





}
