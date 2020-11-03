package main

import "fmt"
func main()  {
	//类型转换
	//bool为不兼容类型，不能转化为其他类型
	var flag bool
	flag = true
	fmt.Printf("%t\n",flag)
	//
	var int1 int8 = 25
	var int2 int64 = 256
	var c1 int64
	c1 =int64(int1) + int2
	fmt.Println(c1)
}
