package main

/*
*
79. 单词搜索
*/
func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int, int) bool
	dfs = func(x int, y int, idx int) (ans bool) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		if vis[x][y] || board[x][y] != word[idx] {
			return false
		}
		vis[x][y] = true
		if idx == len(word)-1 {
			return true
		}
		ans = dfs(x+1, y, idx+1) || dfs(x-1, y, idx+1) || dfs(x, y+1, idx+1) || dfs(x, y-1, idx+1)
		vis[x][y] = false
		return ans
	}
	for i, row := range board {
		for j, x := range row {
			if x == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func main() {

}
