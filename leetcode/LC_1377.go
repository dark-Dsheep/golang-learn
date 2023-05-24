package main

//1377. T 秒后青蛙的位置

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	g := make([][]int, n+1)
	g[1] = []int{0}
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int, int) int
	dfs = func(x, fa, t int) int {
		if t == 0 {
			if x == target {
				return 1
			}
			return 0
		}
		if x == target {
			if len(g[x]) == 1 {
				return 1
			}
			return 0
		}
		for _, y := range g[x] {
			if y != fa {
				prod := dfs(y, x, t-1)
				if prod != 0 {
					return prod * (len(g[x]) - 1)
				}
			}
		}
		return 0
	}
	prod := dfs(1, 0, t)
	if prod == 0 {
		return 0
	}
	return 1 / float64(prod)
}

func main() {

}
