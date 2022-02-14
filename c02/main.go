package main

import "fmt"

func main()  {
	//变量的定义
	//静态语言的变量和动态语言的变量定义差异很大
	//1. 最基础的变量定义
	//var i int
	//i = 20
	//fmt.Println(i)
	//1. 定义并初始化
	var i int = 10
	fmt.Println(i)
	//2. 根据值自行判断变量的类型
	var j = 30
	fmt.Println(j)
	//3. 省略var
	k := 100
	k = 30
	fmt.Println(k)
	//3. 多变量定义
	var a, b, c int = 10, 20, 30
	d, e, f := 10, 20, 30
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	//4. 集合类型
	//var (
	//	a int
	//	name string
	//)

	//定义匿名变量, 变量一旦被定义，

	//常量 -
	const PI = 3.1415926
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	// 常量组如不指定类型和初始变量，该类型和上一行的类型一致
	// 常量的数据类型值可以是布尔， 数字和字符串
	// 不使用的常量， 在编译的时候是不会报错的
	const (
		x int = 16
		y
		s = "abc"
		z
	)

	fmt.Println(x, y, s, z)


	// const 常量的iota, 常量计数器
	// iota 只能在常量组中使用
	// 不同的const定义块互相不干扰
	// 没有表达式的常量定义复用上一行的表达式
	// 从第一行开始， iota从0逐行加1
	const (
		a1 = iota // 计数器
		b1       // 该常量的值等于上一个常量的表达式
		c1
	)

	fmt.Println(a1, b1, c1)
	// 0 1 2

	const (
		unknown = iota
		Female2
		Male2
	)

	fmt.Println(unknown, Female2, Male) // 0 1 2

	const (
		a3 = iota
		b3 = 10
		c3
		d3, e3 = iota, iota
		f3 = iota
	)

	fmt.Println(a3, b3, c3, d3, e3, f3)  // 0 10 10 3


	// 变量的作用域


}
