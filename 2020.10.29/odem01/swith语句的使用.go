package main

import "fmt"

func main()  {
	num :=1
	switch num {
	case 1:
		fmt.Println("按下的一楼")
	case 2:
		fmt.Println("按下的二楼")
	case 3:
		fmt.Println("按下的三楼")
	case 4:
		fmt.Println("按下的四楼")
	default://default表示其他情况
		fmt.Println("按下的是xxx楼")

	}
}
