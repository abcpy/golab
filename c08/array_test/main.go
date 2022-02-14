package main

import "fmt"

func main()  {
	//go语言中的数组和python中的list可以对应起来理解
	//slice和python的list更像
	//静态语言中的数组： 1.大小确定 2. 类型一致
	//数组的声明
	//1. var course [10] string
	//var course = [5] string{"diango", "scrapy", "tornado"}
	course := [5]string{"django", "scrapy", "tornado"}

	// 1. 取值
	fmt.Println(course[0])
	// 2。 修改值
	course[0] = "django3"
	fmt.Println(course[0])

	//数组的另外一种创建方式
	var a = [4]float32{}
	fmt.Println(a)  //[0 0 0 0]

	var c = [5] int{'A', 'B'}
	fmt.Println(c)

	d := [...]int{1,2,3}
	fmt.Println(d)

	//数组操作第一种场景：求长度
	fmt.Println(len(d)) //3
	for i, value := range course {
		fmt.Println(i, value)
	}

	sum := 0
	for _, value := range d {
		sum += value
	}
	fmt.Println(sum)

	//数组是值类型
	courseA := [3]string{"django", "scrapy", "tornado"}
	courseB := [...]string{"django1", "scrapy1", "tornado1","python+go", "asyncio"}

	//go 语言中， courseA 和courseB 都是数组，但是不是同一种类型, 长度不一样类型就不一样
	// go语言中，函数传递参数，数组作为参数，实际调用的是值传递
	fmt.Printf("%T\n", courseA)  //[3]string
	fmt.Printf("%T\n", courseB)  //[5]string






}
