package main

import "fmt"

func main()  {
	//运算符
	var (
		num1 int = 1
		num2 = 2
	)

	ch1 :=num1 + num2
	ch2 :=num1 - num2
	ch3 :=num1 * num2
	ch4 :=num1/num2
	fmt.Println(ch1,ch2,ch3,ch4)
	//非符号运算
	fmt.Println("2 !=3结果为：",2 !=3)
	fmt.Println("!(2>3)结果为：",!(2>3))
	//与符号运算
	fmt.Println("true && true结果为：",true && true)
	fmt.Println("true && false结果为：",true && false)
	fmt.Println("true || true结果为：",true || true)
	//或符号运算
	fmt.Println("true || false结果为：",true || false)
	fmt.Println("false || false结果为：",false || false)

}
