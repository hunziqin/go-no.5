package main

import "fmt"

func main() {
	str := "Go is a beautiful language"
	fmt.Printf("The length of str is: %d\n", len(str))
	for a, b := range str {
		fmt.Printf("Character on position %d is: %c \n", a, b)
	}
	fmt.Println()
	str2 := "Chinese: 中国话"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for a, b := range str2 {
		fmt.Printf("Character %c starts at byte position %d\n", a, b)
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d     %d     %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
