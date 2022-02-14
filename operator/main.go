package main

import "fmt"

func main()  {
	//算术运算符
	a := 12
	b := 22
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	a++
	fmt.Println(a)
	b--
	fmt.Println(b)

	//逻辑运算符 和python不一样
	var c bool = true
	var d bool = false

	if ( c && d ){
		fmt.Printf("第一行 - 条件为true")
	} else {
		fmt.Printf("第一行 - 条件为false\n")

	}

	if ( c || d ){
		fmt.Printf("第一行 - 条件为true\n")
	}

	if !( c && d ){
		fmt.Printf("第一行 - 条件为true\n")
	}

	// 位运算符








}
