package main

import "fmt"

func main()  {
	//给int64起一个别名叫bigint
	type bigint int64
	var a bigint
	fmt.Printf("a type is %T\n",a)
	//类型别名的应用
	type (
		long int64
		char byte
	)
	var a1 long = 11
	var b char = 'a'
	fmt.Printf("a1 = %d,b = %c\n",a1,b)
}
