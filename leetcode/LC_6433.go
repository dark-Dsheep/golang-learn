package main

/*
*
6433. 矩阵中移动的最大次数
*/
var ddd = []struct{ x, y int }{{-1, 1}, {0, 1}, {1, 1}}

func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var bfs func(int, int) int
	bfs = func(v1 int, v2 int) (ans int) {
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		q := [][]int{}
		q = append(q, []int{v1, v2, 0})
		vis[v1][v2] = true
		for len(q) > 0 {
			p := q[0]
			x, y, step := p[0], p[1], p[2]
			q = q[1:] // pop
			ans = max(ans, step)
			for _, d := range ddd {
				nx, ny := x+d.x, y+d.y
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] <= grid[x][y] {
					continue
				}
				if !vis[nx][ny] {
					q = append(q, []int{nx, ny, step + 1})
					vis[nx][ny] = true
				}
			}
		}
		return ans
	}
	for i := 0; i < m; i++ {
		ans = max(ans, bfs(i, 0))
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
