package main

import "fmt"

func main()  {
	//1. make new 函数
	//new 函数用法
	//var p *int
	//*p = 10   //panic: runtime error: invalid memory address or nil pointer dereference

	//默认值 int byte rune float bool string 这些类型都有默认值
	// 指针， slice map interface 默认值是nil

	//对于指针或者说默认值是0的情况来说。如何一开始申明的时候分配内存
	var p *int = new(int) //go的编译器就知道申请一个内存空间，这里的内存中的值全部设置为0
	*p = 10

	//var info map[string]string
	//info["c"] = "jian"  //panic: assignment to entry in nil map

	//除了new函数可以申请内存以外， 还有一个函数是make. 更加常用的是make函数， make 用于slice, map

	var info map[string]string = make(map[string]string)
	info["c"] = "jian"

	//new函数返回的是这个值得地址， 指针， make函数返回的是指定类型的实例

	var info2 map[string]string
	if info2 == nil {
		fmt.Println("map的默认值是nil")
	}

	var slice []string
	if slice == nil {
		fmt.Println("slice的默认值是nil")
	}

	var err error
	if err == nil {
		fmt.Println("error的默认值是nil")
	}





}
