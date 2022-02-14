package main

import "fmt"

type Teacher struct {
	Name string
	Age int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名:%s, 年龄:%d, 职称:%s\n", t.Name, t.Age, t.Title)

}
type Course struct {
	Teacher Teacher //将另一个结构体的变量放进来
	Name string
	Price int
	Url string
}

type Course2 struct {
	Teacher  //将另一个结构体的变量放进来
	Name string
	Price int
	Url string
}

func (c Course) courseInfo()  {
	fmt.Printf("课程名:%s, 价格:%d, 讲师信息: %s %d %s", c.Name, c.Price, c.Teacher.Name, c.Teacher.Age, c.Teacher.Title)
}

func (c Course2) courseInfo2()  {
	fmt.Printf("课程名:%s, 价格:%d, 讲师信息: %s %d %s", c.Name, c.Price, c.Teacher.Name, c.Age, c.Title)
}


func main()  {
   t := Teacher{
   	Name: "jian",
   	Age: 18,
   	Title: "python",
   }

   c := Course{
   	Name: "python",
   	Teacher: t,
   	Url: "",
   }

   fmt.Println(c.Teacher)

	c2 := Course2{
		Name: "python",
		Teacher: t,
		Url: "",
	}

	c2.courseInfo2()


}
