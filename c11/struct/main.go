package main

import (
	"fmt"
	"unsafe"
)


//定义struct 等价于python中的nametuple
type Course struct {
	Name string
	Price int
	Url string
}

func (c Course)test() {
	fmt.Printf("课程名:%s, 课程价格:%d,课程的地址: %s\n", c.Name, c.Price, c.Url)
}
//结构的方法和结构体只能在同一个Package中
//go语言不支持继承， 但是有办法能达到同样的效果

//指针
func (c *Course) setPrice(price int) {
	c.Price = price

}

func main()  {
	//go语言不支持面向对象
	//面向对象的三大特征： 1. 封装 2. 继承 3， 多态 4. 方法重载 5. 抽象基类
	//定义struct go语言没有class这个概念

	//实例化Course k-v形式
	var c Course = Course{
		Name: "django",
		Price: 100,
		Url: "https://xxx.com",
	}

	//访问
	fmt.Println(c.Name, c.Price, c.Url)

	//大小写在go语言中的重要性 可见性

	//一个包中的变量或者结构体如果首字母是小写，那么对于另一个包是不可见的
	//结构体定义的名称 以及属性首字母大写

	//2. 第二种实例化方式 - 顺序形式
	c2 := Course{"flask", 10, "https://www.immoc.com"}
	fmt.Println(c2.Name, c2.Price, c2.Url)

	//3 如果一个指向结构体的指针, 通过结构体指针获取对象的值
	c3 := &Course{"flask", 10, "https://www.immoc.com"}
	fmt.Printf("%T\n", c3) //*main.Course
	//fmt.Println((*c3).Name, (*c3).Price, (*c3).Url)
	fmt.Println(c3.Name, c3.Price, c3.Url)

	//4. 零值 如果不给结构体赋值， go语言会默认给每个字段采用默认值
	c4 := Course{}
	fmt.Println(c4.Name, c4.Price)

	//5. 多种方式零值结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = new(Course)
	//var c8 *Course
	fmt.Println(c5.Price)
	fmt.Println(c6.Price)
	fmt.Println(c7.Price)
	//fmt.Println(c8.Price)//panic: runtime error: invalid memory address or nil pointer dereference
	//指针如果只申明不赋值， 默认值是nil

	// 6. 结构体是值类型
	c8 := Course{"flask", 10, "https://www.immoc.com"}
	c9 := c8
	fmt.Println(c8)
	fmt.Println(c9)
	c8.Price = 200
	fmt.Println(c8)
	fmt.Println(c9)

	//7. 结构体的大小， 占用内存的大小。 可以用sizeof来查看对象占用的类型
	fmt.Println(unsafe.Sizeof(1))  //8
	fmt.Println(unsafe.Sizeof("flask")) //16
	fmt.Println(unsafe.Sizeof(c8))  //40 = 16 + 16 + 8

	//8. slice 的大小
	s1 := []string{"diango", "flask", "pecan"}
	fmt.Println(unsafe.Sizeof(s1))  //24

	m1 := map[string]string{
		"k1":"v1",
		"k2":"v2",
		"k3":"v3",
	}
	fmt.Println(unsafe.Sizeof(m1))  //8

	//结构体方法
	c8.test()
	c8.setPrice(300) //函数参数的传递， 结构体是值传递 用指针改变值
	//结构体接收者有两种形式： 1. 值传递 2. 指针传递，
	c8.test()











}
