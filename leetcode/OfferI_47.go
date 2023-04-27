package main

/**
剑指 Offer 47. 礼物的最大价值
*/

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][210]int, 210)
	f[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		f[0][j] = f[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}
	return f[m-1][n-1]
}

func main() {

}
