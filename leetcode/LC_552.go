package main

/*
*
552. 学生出勤记录 II
*/
func checkRecord(n int) int {
	const mod = 1e9 + 7
	f := make([][2][3]int, n+1)
	for i := 0; i < n+1; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 3; k++ {
				f[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(n int, acnt int, lcnt int) int {
		if acnt >= 2 || lcnt >= 3 {
			return 0
		}
		if n <= 0 {
			return 1
		}
		if f[n][acnt][lcnt] != -1 {
			return f[n][acnt][lcnt]
		}
		ans := 0
		ans = (ans + dfs(n-1, acnt, 0)) % mod
		ans = (ans + dfs(n-1, acnt+1, 0)) % mod
		ans = (ans + dfs(n-1, acnt, lcnt+1)) % mod
		f[n][acnt][lcnt] = ans
		return ans
	}
	return dfs(n, 0, 0)
}

func main() {

}
