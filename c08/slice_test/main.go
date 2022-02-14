package main

import "fmt"

func main()  {
	//什么是切片
	//数组有一个很大的问题：大小确定，不能修改 切片 -- 动态的数组

	//var course [5] string
	//var course []string  //定义了一个切片

	course := []string{"diango", "scrapy", "tornado"}
	fmt.Printf("%T\n", course) //[]string

	//切片的另一种初始化方法 make
	//为什么要传递长度
	course2 := make([]string, 5)
	fmt.Printf("%T\n", course2) //[]string
	fmt.Println(len(course2)) //5

	//通过数组变成一个slice
	//go语言中的切片是一种数据结构
	course3 := [5]string{"diango", "scrapy", "tornado", "python", "asyncio"}
	subCourse := course3[1:4]
	fmt.Printf("%T\n", subCourse) //[]string
	fmt.Println(subCourse) //[scrapy tornado python]

	//new
	course4 := *new([]string)
	fmt.Printf("%T\n", course4) //[]string

	//slice的传递不是值传递 是引用传递

	//slice是动态数组， 需要动态添加值
	fmt.Println(subCourse[0]) //scrapy
	subCourse[0] = "jian"
	fmt.Println(subCourse) //[jian tornado python]

	//append 可以向slice 追加元素
	subCourse = append(subCourse, "xiaoxin")
	fmt.Println(subCourse) //[jian tornado python xiaoxin]

	//拷贝的时候， 目标对象长度需要设置好
	subCourse3 := make([]string, 4)
	copy(subCourse3, subCourse)
	fmt.Println(subCourse3)

	course5 := []string{"diango", "scrapy", "tornado"}
	course6 := append(subCourse, course5...)
	fmt.Println(course6) //[jian tornado python xiaoxin diango scrapy tornado]

	//切片底层是使用数组实现的。既要使用数组又要满足动态的功能。
	// 使用make方法初始化，len和cap是多少
	d := make([]int, 5)
	fmt.Printf("len=%d, cap=%d", len(d), cap(d)) //len=5, cap=5

	d1 := make([]int, 5, 10)
	fmt.Printf("len=%d, cap=%d\n", len(d1), cap(d1)) //len=5, cap=10

	//通过数组取切片
	e := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	e_slice := e[2:4]
	fmt.Printf("len=%d, cap=%d", len(e_slice), cap(e_slice)) //len=2, cap=8

	//当make遇到了append的时候， 容易遇到坑


























}
