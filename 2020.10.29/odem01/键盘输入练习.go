package main

import "fmt"
func main() {
	str := "abcdefgh"
	fmt.Println(str)
	var a int
	fmt.Println("请输入你想打打印第几个字母(0-7)")
	fmt.Scan(&a)
	fmt.Printf("你想打印的字母是%c (小写)\n",str[a])
	var b byte
	b = 65
	c:= b+byte(a)
	fmt.Printf("你想打印的字母是%c (大写)\n", c)
}
