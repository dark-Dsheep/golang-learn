package main

/*
*
1245. 树的直径
*/
func treeDiameter(edges [][]int) (ans int) {
	n := len(edges)
	g := make([][]int, n+1)
	for _, e := range edges {
		g[e[1]] = append(g[e[1]], e[0])
		g[e[0]] = append(g[e[0]], e[1])
	}
	var dfs func(int, int) int
	dfs = func(x int, fa int) int {
		v1 := 0
		for _, y := range g[x] {
			if y != fa {
				v2 := dfs(y, x) + 1
				ans = max(ans, v1+v2)
				v1 = max(v1, v2)
			}
		}
		return v1
	}
	dfs(0, -1)
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {
	edges := [][]int{{0, 1}, {0, 2}}
	treeDiameter(edges)
}
