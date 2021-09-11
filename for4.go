package main

func main() {
	a := 0
HERE:
	println(a)
	a++
	if a == 15 {
		return
	}
	goto HERE
}
