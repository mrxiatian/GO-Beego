package main

import "fmt"

func main()  {
	var i int =112
	if i == 12 {
		fmt.Println("i为10")
	}else if i<10 {
		fmt.Println("i的值小于是0")
	}else if i>10 {
		fmt.Println("i的值大于10")
	}else {
		fmt.Println("error")
	}

	var num int
	fmt.Println("请输入你想按的数值")
	fmt.Scan(&num)
	switch  {
	case num < 0:
		fmt.Println("1")
	case num == 0:
		fmt.Println("2")
	case num > 0:
		fmt.Println("3")



	}


}
