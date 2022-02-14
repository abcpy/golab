package main

import (
	"errors"
	"fmt"
)

//函数的几个要素： 1. 函数名 2. 参数 3. 返回值

//函数第一种定义
func add(a,b int) int {
	return a + b
}

//函数的第二种定义方法
func add2(a,b int) (sum int)  {
	sum = a + b
	return sum
}

// 函数的第三种定义方法
func add3(a, b int) (sum int) {
	sum = a + b
	return
}

// 函数的第四种定义方法
// 被除数等于0， 要返回多个值，
func div(a,b int) (int, error) {
	var result int
	var err error
	if b == 0 {
		err = errors.New("被除数不能为0")
	}else {
		result = a / b
	}

	return result, err
}

// 被除数等于0， 要返回多个值，
func div2(a,b int) (result int, err error) {
	if b == 0 {
		err = errors.New("被除数不能为0")
	}else {
		result = a / b
	}

	return
}

func main()  {
	result, err := div(12, 0)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	result2, err2 := div2(12, 3)
	if err2 != nil {
		fmt.Println(err2)
	}else {
		fmt.Println(result2)
	}
}
