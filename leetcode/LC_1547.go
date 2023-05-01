package main

import "sort"

/**
1547. 切棍子的最小成本
*/

func minCost(n int, cuts []int) int {
	f := make([][110]int, 110)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	sort.Ints(cuts)
	cuts = append([]int{0}, cuts...)
	cuts = append(cuts, n)
	var dfs func(int, int) int
	dfs = func(i int, j int) (ans int) {
		if i+1 >= j {
			return 0
		}
		ptr := &f[i][j]
		if *ptr != -1 {
			return *ptr
		}
		ans = 0x3f3f3f3f
		for k := i + 1; k < j; k++ {
			v := cuts[j] - cuts[i]
			ans = min(ans, v+dfs(i, k)+dfs(k, j))
		}
		*ptr = ans
		return ans
	}
	return dfs(0, len(cuts)-1)
}

func main() {

}
