package main

/**
6419. 使二叉树所有路径值相等的最小代价
*/

func minIncrements(n int, cost []int) (ans int) {
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		l := dfs(i*2 + 1)
		r := dfs(i*2 + 2)
		ans += max(l, r) - min(l, r)
		cost[i] += max(l, r)
		return cost[i]
	}
	dfs(0)
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
