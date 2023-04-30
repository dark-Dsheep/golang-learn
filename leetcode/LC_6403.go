package main

/*
*

6403. 网格图中鱼的最大数目
*/

var dirs = []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x int, y int) (ans int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		ans = grid[x][y]
		grid[x][y] = 0
		for _, d := range dirs {
			ans += dfs(x+d.x, y+d.y)
		}
		return ans
	}
	for i, row := range grid {
		for j := range row {
			if grid[i][j] != 0 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

func main() {

}
