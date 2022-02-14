package main

import "fmt"

func main()  {
	/*
	defer 语句是go中的一种用于注册延迟调用的机制， 它可以让当前函数执行完毕之后执行
	defer之后只能是函数调用， 不能是表达式 比如 a++
	test1
	test3
	defer test2
	多个defer是按照先入后出的顺序执行的
	*/
	//fmt.Println("test1")
	//defer fmt.Println("defer test2")
	//defer fmt.Println("defer test4")
	//defer fmt.Println("defer test5")

	//fmt.Println("test3")

	// defer语句执行时的拷贝机制
	test := func() {
		fmt.Println("test1")
	}
	defer test()

	test = func() {
		fmt.Println("test2")
	}

}
