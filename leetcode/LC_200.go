package main

/*
*
200. 岛屿数量
*/

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x int, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}
	for i, row := range grid {
		for j, v := range row {
			if v == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

func main() {

}
