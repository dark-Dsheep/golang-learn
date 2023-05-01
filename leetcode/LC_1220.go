package main

/*
*
1220. 统计元音字母序列的数目
*/
func countVowelPermutation(n int) (ans int) {
	cs := []int{'a', 'e', 'i', 'o', 'u'}
	f := make([][5]int, 20001)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(n int, i int) (ans int) {
		if n == 0 {
			return 1
		}
		ptr := &f[n][i]
		if *ptr != -1 {
			return *ptr
		}
		c := cs[i]
		if c == 'a' {
			ans = (ans + dfs(n-1, 4)) % 1000000007
			ans = (ans + dfs(n-1, 1)) % 1000000007
			ans = (ans + dfs(n-1, 2)) % 1000000007
		} else if c == 'e' {
			ans = (ans + dfs(n-1, 0)) % 1000000007
			ans = (ans + dfs(n-1, 2)) % 1000000007
		} else if c == 'i' {
			ans = (ans + dfs(n-1, 1)) % 1000000007
			ans = (ans + dfs(n-1, 3)) % 1000000007
		} else if c == 'o' {
			ans = (ans + dfs(n-1, 2)) % 1000000007
		} else if c == 'u' {
			ans = (ans + dfs(n-1, 2)) % 1000000007
			ans = (ans + dfs(n-1, 3)) % 1000000007
		}
		*ptr = ans
		return ans
	}

	for i := 0; i <= 4; i++ {
		ans = (ans + dfs(n-1, i)) % 1000000007
	}
	return ans
}

func main() {

}
