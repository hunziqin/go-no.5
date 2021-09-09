package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "ABC"
	var b string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(a)

	if err != nil {
		fmt.Printf("a %s is not an integer - exiting with error\n", a)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an += 5
	b = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", b)
	return
}
