package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Println(j)
		}
		fmt.Println(" ")
	}
}
