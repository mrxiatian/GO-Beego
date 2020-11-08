package main

import "fmt"

func main() {
	str := "abc"
	for i :=0; i<len(str); i++{
		fmt.Printf("str[%d] = %c\n", i , str[i])
		for i,bata :=range str{
			fmt.Printf("str[%d] = %c\n", i , bata)
		}
		for i :=range str{
			fmt.Printf("str[%d] = %c\n", i,str[i])
		}
	}
}
