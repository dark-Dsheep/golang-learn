package main

import "fmt"

/**
2652. 倍数求和
*/

func sumOfMultiples(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			ans += i
		}
	}
	return ans
}

func main() {

	ans := sumOfMultiples(7)
	fmt.Println(ans)
}
