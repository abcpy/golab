package main

import (
	"fmt"
)

func main()  {
	//go语言中的map --> python中的dict
	// go语言中的map的key和value 类型申明时就要指明

	m1 := map[string]string{
		"m1":"v1",
	}
	fmt.Printf("%v\n", m1) //map[m1:v1]

	//make 函数创建
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2) //map[m2:v2]

	//定义一个空的map
	m3 := map[string]string{}
	fmt.Printf("%v\n", m3) //map[]

	//map中的key不是所有的类型都支持，该类型需要支持 == 或者 != 操作
	//map的基本操作
	m4 := map[string]string{
		"a": "va",
		"b": "vb",
	}

	m4["a"] = "vc"
	m4["b"] = "vd"
	fmt.Printf("%v\n", m4) //map[a:vc b:vd]

	//查询
	fmt.Println(m4["a"]) //vc

	fmt.Println(m4["e"])

	delete(m4, "a")

	//遍历
	for k,v := range m4{
		fmt.Println(k,v)
	}












}
