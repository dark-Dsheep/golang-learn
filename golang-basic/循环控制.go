package main

import "fmt"

func main() {

	//for i := 0; i < 20; i++ {
	//	fmt.Println(i)
	//}
	//
	//for i, j := 1, 2; i < 100 && j < 1000; i++ {
	//	fmt.Println(i, j)
	//}
	//
	//num := 1
	//for num < 100 {
	//	num *= 2
	//}

	// for range
	s := "hello world"
	for idx, c := range s {
		fmt.Println(idx, c)
	}
}
