package main

/**
312. 戳气球
*/

func maxCoins(nums []int) int {
	n := len(nums)
	a := make([]int, n+2)
	a[0], a[n+1] = 1, 1
	for i, x := range nums {
		a[i+1] = x
	}
	f := make([][]int, n+2)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+2)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i+1 >= j {
			return 0
		}
		if f[i][j] != -1 {
			return f[i][j]
		}
		ans := -0x3f3f3f
		for k := i + 1; k < j; k++ {
			v := a[i] * a[k] * a[j]
			ans = max(ans, v+dfs(i, k)+dfs(k, j))
		}
		f[i][j] = ans
		return ans
	}
	return dfs(0, len(a)-1)
}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}

func main() {

}
