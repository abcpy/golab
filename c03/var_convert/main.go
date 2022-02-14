package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1. 基本类型的转换
	a := int(3.0)
	fmt.Println(a)
	//2. go语言中不支持变量间的隐式类型转换
	//变量间类型转换不支持
	var b int = 5.0  //5.0 是常量  常量到变量支持隐式转换
	fmt.Println(b)

	//var b2 int = 5.1  //5.1' (type untyped float) cannot be represented by the type in
	//fmt.Println(b2)

	c := 6.0
	fmt.Printf("%T\n", c)  //float64

	//var d int = c  //Cannot use 'c' (type float64) as the type int

	var d int = int(c)
	fmt.Println(d)

	c2 := 6.1
	var d2 int = int(c2)
	fmt.Println(d2)  //6

	//var d3 int = int("1") //Cannot convert an expression of the type 'string' to the type 'int'

	var e int32 = 56
	var f int64 = int64(e)
	fmt.Println(f)

	// int 转字符串 itoa
	var g int = 10
	s := strconv.Itoa(g)
	fmt.Printf("%T, %v\n", s, s) //string, 10

	// 字符串转int aoti
	data, _ := strconv.Atoi("45")
	fmt.Printf("%T, %v\n", data, data)

	//parse 类函数
	m, _ := strconv.ParseBool("False")
	fmt.Printf("%T, %v\n", m, m)

	m2, err := strconv.ParseBool("H")
	fmt.Printf("%v, %v\n", m2, err) //false, strconv.ParseBool: parsing "H": invalid syntax

	m3, _ := strconv.ParseFloat("3.1415", 32)
	fmt.Printf("%T, %v\n", m3, m3) //float64, 3.1414999961853027

	m4, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("%T, %v\n", m4, m4) //float64, 3.1415

	m5, _ := strconv.ParseInt("20",10,32)
	fmt.Printf("%T, %v\n", m5, m5) //int64, 20

	m6, _ := strconv.ParseInt("20",10,64)
	fmt.Printf("%T, %v\n", m6, m6) //int64, 20

	m7, _ :=strconv.ParseUint("42", 10, 64)
	fmt.Printf("%T, %v\n", m7, m7) //uint64, 42

	m8, _ :=strconv.ParseUint("42", 10, 32)
	fmt.Printf("%T, %v\n", m8, m8) //uint64, 42

	//其他类型转换字符串
	 m9 := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", m9, m9) //string, true

	m10 := strconv.FormatInt(-42, 10)
	fmt.Printf("%T, %v\n", m10, m10) //string, -42

	//  python的类型相比go而言少很多


















}
