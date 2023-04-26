package main

/**
509. 斐波那契数
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	c := a + b
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}

func main() {

}
