package main

import (
	"fmt"
)

//接口是一个协议 -- 一组方法的集合
type Programmer interface {
	//定义方法 申明
	Coding() string
	Debug() string
}

type Desginer interface {
	Desgin() string
}

type Manager interface {
	Programmer
	Desginer
	Manage() string
}



type UIDesginer struct {

}

func (ui UIDesginer)Desgin() string  {
     fmt.Println("我会设计")
     return "我会ui设计"
}

type Pythoner struct {
	UIDesginer

}


func (p Pythoner) Coding() string {
	fmt.Println("python开发者")
	return "Python 开发者"
}

func (p Pythoner) Debug() string {
	fmt.Println("我会Debug")
	return "Python debug 开发者"
}

//func (p Pythoner) Desgin() string {
//	fmt.Println("我会Desiner")
//	return "Desginer开发者"
//}

func (p Pythoner) Manage() string {
	fmt.Println("manage")
	return "manage"
}



//多态: go语言中interface是一种类型， 是一种抽象类型



//申明的类型是一种兼容类型，但是实际赋值的时候是另一种类型
//python不需要多态，因为python是动态的

func main()  {
	var pro Programmer = Pythoner{}
	pro.Coding()

	//go struct 组合
	//组合一起实现了所有接口的方法

    var m Manager = Pythoner{}
    fmt.Println(m)

    //var err error = errors.New("错误")

    s := "文件不存在"
    var err error = fmt.Errorf("%s", s)
    fmt.Println(err)



}
