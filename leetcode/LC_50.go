package main

import "fmt"

/**
50. Pow(x, n)
*/

func myPow(x float64, n int) float64 {
	return fastPow(x, n)
}

// 快速幂
func fastPow(m float64, k int) float64 {
	if k < 0 {
		k = -k
		m = 1 / m
	}
	ans, t := 1.0, m
	for k > 0 {
		if (k & 1) == 1 {
			ans = ans * t
		}
		t = t * t
		k >>= 1
	}
	return ans
}

func main() {
	ans := myPow(2.00000, 10)
	fmt.Println(ans)
}
