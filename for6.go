package main

import "fmt"

func main() {
	for a := 1; a < 101; a++ {
		switch {
		case a%3 == 0:
			fmt.Println("Fizz")
		case a%5 == 0:
			fmt.Println("Buzz")
		case a%3 == 0 && a%5 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(a)
		}
	}
}
