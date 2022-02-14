package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//结构体标签 omitempty 标签在序列化的时候可以忽略0或空值
type Info struct {
	Name string `json:"name"`
	AgeDetail int `json:"age,omitempty"`
	Gender string `json:"-"`
}

//反射包
func main()  {
	info := Info{
		Name: "jian",
		Gender: "Male",
	}

	re, _ := json.Marshal(info)
	fmt.Println(string(re))

	//通过反射包识别这些tag
	t := reflect.TypeOf(info)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		tag := filed.Tag.Get("json")
		fmt.Println(tag)
	}


	//接口 (protocol) 参考了python中的鸭子类型
	//

}
