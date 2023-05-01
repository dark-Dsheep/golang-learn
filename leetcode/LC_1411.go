package main

/*
*
1411. 给 N x 3 网格图涂色的方案数
*/
func numOfWays(n int) int {

	f := make([][2]int64, 5001)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int64
	dfs = func(i int, c int) (ans int64) {
		if i >= n {
			return 1
		}
		if i == 0 {
			return 6 * (dfs(i+1, 0) + dfs(i+1, 1)) % 1000000007
		}
		ptr := &f[i][c]
		if *ptr != -1 {
			return *ptr
		}
		if c == 1 {
			ans = (ans + 2*dfs(i+1, 0) + 2*dfs(i+1, 1)) % 1000000007
		} else {
			ans = (ans + 3*dfs(i+1, 0) + 2*dfs(i+1, 1)) % 1000000007
		}
		*ptr = ans
		return ans
	}
	return int(dfs(0, 0))
}

func main() {

}
