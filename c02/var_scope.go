package main

import "fmt"

var c = 20

// python没有定义变量的说法  全局变量 global


func main()  {

	// 变量的作用域
    // 局部变量
	//c := 10
	var c int = 10
	fmt.Println(c)

}
