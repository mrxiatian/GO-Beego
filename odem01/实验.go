package main

import "fmt"

func main()  {
	var a float32 = 3.13
	fmt.Println(a)
	b,_ :=1,3
	fmt.Println(b)
	var str1,str2 = "gsgs","sgdsag"
	fmt.Println(str1,str2)
	var (
		name = "xiatian"
		age = 2
	)
	fmt.Printf("名字是：%s,年龄为：%d",name,age)
	const PAT float32  = 3.12
	const BAIDU  = "http://www.baidu.com"
	fmt.Println(PAT)
	fmt.Println(BAIDU)
	const (
		const1 int = 1
		const2 string = "sjag"
		const3 bool = true
		const4 byte = 'a'
	)
	fmt.Printf("%d,%s,%t,%c\n",const1,const2,const3,const4)
	const (
		a1 = iota
		a2 ,a3 ,a4=iota,iota,iota
		a5 = iota
	)
	fmt.Println(a1,a2,a3,a4,a5)
	var ch byte
	ch = 97
	fmt.Printf("%c,%d\n",ch)

}
