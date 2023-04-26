package main

import "math"

/*
*
1039. 多边形三角剖分的最低得分
*/
func minScoreTriangulation(values []int) int {
	f := make([][60]int, 60)
	for i := range f {
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (ans int) {
		ptr := &f[i][j]
		if *ptr != -1 {
			return *ptr
		}
		if i+1 >= j {
			return
		}
		ans = math.MaxInt
		for k := i + 1; k < j; k++ {
			v := values[i] * values[k] * values[j]
			ans = min(ans, v+dfs(i, k)+dfs(k, j))
		}
		*ptr = ans
		return ans
	}
	return dfs(0, len(values)-1)
}

func main() {

}
