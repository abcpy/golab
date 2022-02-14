package main

import "fmt"

func main() {
	//go 语言中的关键词 type
	//1. 给一个类型定义别名 为了强调现在处理的对象是字节类型
	//2. 这种别名实际上是为了代码的可读性。 实际本质上还是uint8, 就是在代码编译阶段可读性强
	//var a byte
	type myByte = byte
	var b myByte
	fmt.Printf("%T\n", b) // unit8

	//2. 基于一个已有的类型定义一个新的类型
	type myInt int
	var i myInt
	fmt.Printf("%T", i) // main.myInt

	//3. 定义结构体
	type Course struct {

	}
	//4. 接口
	type Callable interface {

	}

	type handle func(str string)


}
