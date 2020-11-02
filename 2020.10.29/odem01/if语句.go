package main

import (
	"fmt"
)

func main()  {
	//if语句基本
	s :=110
	if s == 110 {
		fmt.Println("s的值为110")
	}
	var a int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Println("请输入你正确的年龄")
	}
	if a>0 && a<18 {
		fmt.Println("你还是个孩子，还不能结婚")
	}
	if a == 18 {
		fmt.Println("恭喜你成年了，可以结婚了")
	}
	if a >18 && a <=120 {
		fmt.Println("祝你生活愉快")
	}
	if a >120 {
		fmt.Println("你是妖怪吗？")
	}
	//if语句进阶,else的使用
	a1 :=29
	if a1 ==29 {
		fmt.Println("a1 = 29")
	}else if a1 >29 {
		fmt.Println("a1大于29")
	}else if a1 <29 {
		fmt.Println("a1小于29")
	}
}
