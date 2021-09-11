package main

import "fmt"

func main() {
	var g string = "G"
	for a := 0; a < 25; a++ {
		for b := 0; b <= a; b++ {
			fmt.Printf("%c", g)
		}
	}
}
