package main

import "fmt"

func main() {
	var a int = 7

	switch {
	case a < 0:
		fmt.Println("Number is negative")
	case a > 0:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}
