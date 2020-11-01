package main

import "fmt"

func main()  {
	a := 10
	b := "anc"
	c := 'a'
	d := 3.14
	e := true
	//%T操作变量所属类型
	fmt.Printf("a = %T\n,b = %s\n,c = %c\n, d = %f\n,e = %t\n",a,b,c,d,e)

	//输入的使用
	var a1 int
	fmt.Printf("请输入数值：")

	//阻塞等待用户的输入
	fmt.Scan(&a1)
	fmt.Println("变量的值为:",a1)



}
