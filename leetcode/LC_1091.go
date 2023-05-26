package main

//1091. 二进制矩阵中的最短路径

//var dirs = []struct{ x, y int }{{0, 1}, {1, 0}, {1, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	q := [][3]int{{0, 0, 1}}
	for len(q) > 0 {
		p := q[0]
		x, y, step := p[0], p[1], p[2]
		if x == m-1 && y == n-1 {
			return step
		}
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] != 1 {
				grid[nx][ny] = 1
				q = append(q, [3]int{nx, ny, step + 1})
			}
		}
	}
	return -1
}

func main() {

}
