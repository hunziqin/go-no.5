package main

import "fmt"

func main() {
	var a = 10
	var b int

	if a <= 0 {
		fmt.Println("a<0 or a=0")
	} else if a > 0 && a < 5 {
		fmt.Println("a>0 and a<10")
	} else {
		fmt.Println("a=10")
	}

	if b = 5; b > 10 {
		fmt.Printf("b is greater than 10\n")
	} else {
		fmt.Printf("b is not greater than 10\n")
	}
}
