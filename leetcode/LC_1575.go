package main

/*
*
1575. 统计所有可行路径
*/
const MOD = 1e9 + 7

func countRoutes(loc []int, start int, finish int, fuel int) int {
	f := make([][210]int, 110)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(pos int, tot int) (ans int) {
		if tot < 0 {
			return 0
		}
		ptr := &f[pos][tot]
		if *ptr != -1 {
			return *ptr
		}
		if pos == finish {
			ans = 1
		}
		for i, v := range loc {
			if i == pos {
				continue
			}
			ans += dfs(i, tot-abs(loc[pos]-v))
			ans %= MOD
		}
		*ptr = ans
		return ans
	}
	return dfs(start, fuel)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
