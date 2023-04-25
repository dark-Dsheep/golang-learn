package main

import "fmt"

func m1() {
	a, b := 1, 2
	if a > b {
		b++
	} else {
		a++
	}
}

func m2() {
	if x := 1 + 1; x > 2 {
		fmt.Println(x)
	}
}

func main() {
	m1()
	m2()
}
