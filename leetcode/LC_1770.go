package main

import "math"

/*
*
1770. 执行乘法运算的最大分数
*/
func maximumScore(a []int, b []int) int {
	f := make([][1010]int, 1010)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i int, j int, idx int) (ans int) {
		if idx >= len(b) {
			return 0
		}
		ptr := &f[i][len(a)-j]
		if *ptr != -1 {
			return *ptr
		}
		ans = math.MinInt
		ans = max(ans, a[i]*b[idx]+dfs(i+1, j, idx+1))
		ans = max(ans, a[j]*b[idx]+dfs(i, j-1, idx+1))
		*ptr = ans
		return ans
	}
	return dfs(0, len(a)-1, 0)

}

func main() {

}
