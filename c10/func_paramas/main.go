package main

import "fmt"

//省略号 通过省略号去动态设置多个参数值
func add(params ...int) (sum int) {
    for _, v := range params {
    	sum += v
	}
	return
}


func main()  {
	fmt.Println(add(1, 2)) //3
	fmt.Println(add(1, 2 ,3)) //6

	slice := []int{1, 2, 3, 4, 5} //15
	fmt.Println(add(slice...))

	//省略号的用途 1. 函数参数不定长 2. 将slice打散
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr) //[3]int

	//go语言中非常重要的特性 函数  一等公民 可以作为参数 返回值 赋值给一个变量
	myfunc := add
	fmt.Printf("%T\n", myfunc) //func(...int) int

	//匿名函数
	result := func (a, b int) int {
		return a + b
	}

	fmt.Println(result(1, 2)) //3

	//python中的finally是在return 之前调用的
	//finally中是可以return的， 会覆盖原有的return

	//go语言并没有提供try... except...finally





}
