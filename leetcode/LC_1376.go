package main

/*
*
1376. 通知所有员工所需的时间
*/
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i, fa := range manager {
		if i == headID {
			continue
		}
		g[fa] = append(g[fa], i)
	}
	var dfs func(int) int
	dfs = func(x int) (ans int) {
		for _, y := range g[x] {
			ans = max(ans, informTime[y]+dfs(y))
		}
		return ans
	}
	return dfs(headID) + informTime[headID]
}

func main() {

}
