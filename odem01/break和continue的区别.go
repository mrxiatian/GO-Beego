package main

import (
	"fmt"
	"time"
)

func main() {
	for i :=0;i<10; i++{
		time.Sleep(time.Second)
		if i ==5 {
			continue
		}
		if i == 8 {
			break
		}
		fmt.Println("i = ",i)
	}
}
