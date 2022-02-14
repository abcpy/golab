package main

import (
	"fmt"
)

func main()  {
	//for
	//for init; condition; post{}
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// for condition {}
	//j := 0
	//for j <=10 {
	//	fmt.Println(j)
	//	j++

	name := "jianliu"
	//for index, value := range name {
	//	fmt.Println(index, value)
	//}

	for _, value := range name {
		fmt.Printf("%c\n",value)
	}

	for _, value := range name {
		fmt.Println(string(value))
	}

	name_arr := []rune(name)  //type rune = int32
	fmt.Println(len(name_arr))
	for i := 0; i < len(name_arr); i++ {
		fmt.Printf("%c\n", name_arr[i])

	}

	}
	// for {}
