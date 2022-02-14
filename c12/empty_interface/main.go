package main

import "fmt"

//type EmptyInterface interface {
//
//}

type course struct {
	name string
	price int
	url string
}

func print(x interface{})  {
	switch v := x.(type) {
	case int:
		fmt.Printf("%d(整数)", v)
	}
	//v, ok := x.(int)
	////go的默认问题
	//if ok {
	//	//x是int类型
	//	fmt.Printf("%d, 整数", v)
	//}
	//fmt.Printf("%v\n", i)

}
func main()  {
	//空接口

	var i interface{}
	//空接口类似于python中的object

	i = course{}
	fmt.Println(i)
	var a = 10
	print(a)

	var b = []string{"python", "Java"}
	print(b)

	//1. 空接口的第一个用途， 可以把任何类型都赋值给空接口变量
	//2. 参数传递
	//3. 空接口可以作为map的值
	//var teacherInfo = make(map[string]string)

	var teacherInfo = make(map[string]interface{})
	teacherInfo["k1"] = "v1"
	teacherInfo["k2"] = 12
	teacherInfo["k3"] = 13
	fmt.Println(teacherInfo)
	fmt.Printf("%v", teacherInfo)

	//类型断言




}
