package main

/*
*
463. 岛屿的周长
*/
func islandPerimeter(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 1 {
			grid[x][y] = -1
			dfs(x+1, y)
			dfs(x-1, y)
			dfs(x, y+1)
			dfs(x, y-1)
		}
	}
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	return ans
}

func main() {

}
