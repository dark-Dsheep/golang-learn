package main

import "fmt"

/**
70. 爬楼梯
*/

func climbStairs(n int) int {
	a, b, c := 1, 1, 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {

	ans := climbStairs(2)
	fmt.Println(ans)
}
