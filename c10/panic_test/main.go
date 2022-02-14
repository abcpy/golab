package main

import "strconv"

func main()  {
	 //错误就是能遇到可能出现的情况
	_, err := strconv.Atoi("abcd")
	if err != nil {
		//错误
	}

	//异常 go语言中如何抛出异常和捕捉异常
	//panic会引起主线程的挂掉。 同时会导致其他的协程都挂掉
	//recover() panic()

}
