package main

import (
	"fmt"
	"strings"
)

func main()  {
	// 字符串基本操作
	// 1. 求解字符串的长度
	//涉及到了中文字符会发生变化

	var name string = "liujian:刘健"
	fmt.Println(len(name)) //14

	name_arr := []rune(name)
	fmt.Println(len(name_arr)) //10

	name_arr32 := []int32(name)
	fmt.Println(len(name_arr32)) //10

	date := `2021/12/25`
	fmt.Println(date)

	//是否包含某个子串
	fmt.Println(strings.Contains(name, "liu")) //true
	fmt.Println(strings.Index(name, "liu"))   //0
	fmt.Println(strings.Index(name, "jian"))   //3

	// 统计出现的次数
	fmt.Println(strings.Count(name, "i"))  //2

	//前缀和后缀
	fmt.Println(strings.HasPrefix(name, "l")) //true
	fmt.Println(strings.HasSuffix(name, "健")) //true

	//大小写转换
	fmt.Println(strings.ToUpper("liujian")) //LIUJIAN
	fmt.Println(strings.ToLower("LIUJIAN")) //liujian

	//字符串比较
	fmt.Println(strings.Compare("a", "b")) //-1
	fmt.Println(strings.Compare("a", "a")) //0
	fmt.Println(strings.Compare("b", "a")) //1
	fmt.Println(strings.Compare("a", "10")) //1

	//去掉空格和指定字符串
	fmt.Println(strings.TrimSpace("liujian "))
	fmt.Println(strings.TrimLeft("liujian", "l")) //iujian
	fmt.Println(strings.TrimRight("liujian", "n")) //liujia

	//split 方法
	fmt.Println(strings.Split("a,b,c",","))  //[a b c]

	//合并join方法将字符串数组连接起来
	s := []string{"foot", "bar", "baza"}
	fmt.Println(strings.Join(s, "-"))  //foot-bar-baza

	//字符串替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) //oinky oinky oink












}
