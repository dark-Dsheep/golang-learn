package main

/*
*2246. 相邻字符不同的最长路径
 */
func longestPath(parent []int, s string) (ans int) {
	n := len(parent)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		fa := parent[i]
		g[fa] = append(g[fa], i)
	}
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		for _, y := range g[x] {
			len := dfs(y) + 1
			if s[y] != s[x] {
				ans = max(ans, maxLen+len)
				maxLen = max(maxLen, len)
			}
		}
		return maxLen
	}
	dfs(0)
	return ans + 1
}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}

func main() {

}
