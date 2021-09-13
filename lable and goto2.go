package main

import "fmt"

func main() {
	i := 0
	for { //由于没有检查，这是一个无限循环
		if i <= 3 {
			break
		} //满足此条件时跳出此 for 循环
		fmt.Printf("Value of i is: %d", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}
