package main

import "math"

/*
*
1000. 合并石头的最低成本
*/
func mergeStones(stones []int, k int) int {
	f := make([][40][40]int, 40)
	n := len(stones)
	if (n-1)%(k-1) != 0 {
		return -1
	}
	sum := make([]int, n+1)
	for i, x := range stones {
		sum[i+1] = sum[i] + x
	}
	for i := range f {
		for j := range f[i] {
			for k := range f[i][j] {
				f[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, c int) (ans int) {
		ptr := &f[i][j][c]
		if *ptr != -1 {
			return *ptr
		}
		defer func() { *ptr = ans }()
		if c == 1 {
			if i == j {
				return
			}
			return dfs(i, j, k) + sum[j+1] - sum[i]
		}
		ans = math.MaxInt
		for m := i; m < j; m += k - 1 {
			ans = min(ans, dfs(i, m, 1)+dfs(m+1, j, c-1))
		}
		return ans
	}
	return dfs(0, n-1, 1)
}

func main() {

}
