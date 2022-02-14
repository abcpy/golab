package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main()  {
	// 定义bool 类型
	var gender bool = false
	gender2 := true
	fmt.Println(gender, gender2)
	// 相比pyton而言， go语言为什么有这么多种整数类型
	// 年龄， 分数都是有上线的
	// 很多场景之下， 数字有上限， 我们可以选择

	//int 是一种动态类型，取决于机器本身是多少位， 64位机器上运行那么int, 就是int64,
	// 如果32位机器上那么就是4个字节
	//一般情况下我们都会指明int占用多少个字节
	var age int8 = 18
	var age2 int = 20
	fmt.Println(age, age2)
	fmt.Println(unsafe.Sizeof(age2))  //8个字节


	// float 类型 float类型最大数
	// float32 和 float64 两者占用内存不一样，
	var weight float32 = 71.2
	fmt.Println(weight)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	weight2 := 78.4
	fmt.Printf("%T", weight2) //float64
	fmt.Println(unsafe.Sizeof(weight2))// 48

	// byte rune
	//byte 类型
	var a byte = 18
	fmt.Println(a)

	// a 可以和数字计算
	// a + 1的类型是int32
	// int类型可以直接变成字符
	b := 'b'
	fmt.Printf("%T\n", b + 1) //int32
	fmt.Printf("a+1=%c", b+1)



}
