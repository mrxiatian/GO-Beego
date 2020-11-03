package main

import (
	"fmt"
)

func main()  {
	var a float32=3.1433333333333333333
	fmt.Println(a)
	//_的应用
	b, _ :=1, 3
	fmt.Println(b)

	//多重赋值

	var str1, str2 = "x99a","diig"
	fmt.Println(str1,str2)
	var (
		name = "xiatian"
		age = 20
	)
	fmt.Printf("名字是：%s，年龄为：%d",name,age)
	//常量
	const PAT float32  = 3.14
	const BAIDU = "http://www.baidu.com"
	fmt.Println(PAT)
	fmt.Println(BAIDU)

	const (
		x int = 10
		z int = 23
		s
		j
		name1 string = "xiaosan"
	)
	fmt.Println(x,z,s,j,name1)
	//iota枚举的应用
	const (
		a1 = iota
		a2,a3,a4 = iota,iota,iota
		a5 = iota
	)
	fmt.Println(a1,a2,a3,a4,a5)
	//布尔类型
	var b4 bool
	fmt.Println(b4)

	var b1 = false
	fmt.Println(b1)

	b2 := false
	fmt.Println(b2)
	//字符类型
	//格式化输出，%c以字符方式打印，%d的以整数方式打印
	var ch byte
	ch = 97
	fmt.Printf("%c,%d\n",ch,ch)

	ch = 'a'
	fmt.Printf("%c,%d\n",ch,ch)

	//大小写转换
	a6 :='A'
	fmt.Printf("大写：%d, 小写：%d\n",'A','a')
	fmt.Printf("大写转小写： %c\n",a6+32)
	fmt.Printf( "小写转大写： %c\n",a6-32)

	//len（）的应用
	b3 := "string"
	fmt.Println("len(b3)的值为：",len(b3))

//字符串类型
	var str3 string
	str3 = "abc"
	fmt.Println("str3 = ",str3)

	//自动推到类型
	str4 := "mike"
	fmt.Printf("str4 类型是 %T\n",str4)

	//内建函数，len（）可以测试字符串的长度，有多少个字符
	fmt.Println("len(str4) = ", len(str4))

	//只想操作字符串的某个字符，从开始操作
	fmt.Printf("str4[0] = %c, str4[1] = %c\n",str4[0],str4[1])

	//复数的类型
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ",t)
	t2 := 3.3+4.4i
	fmt.Printf("t2 type is %T\n",t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ",real(t2),"imag(t2) = ",imag(t2))


}
