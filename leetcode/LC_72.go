package main

/**
72. 编辑距离
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	word1 = " " + word1
	word2 = " " + word2
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		f[i][0] = i
	}
	for j := 1; j <= n; j++ {
		f[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i] == word2[j] {
				f[i][j] = f[i-1][j-1]
			} else {
				f[i][j] = min(f[i-1][j-1], min(f[i-1][j], f[i][j-1])) + 1
			}
		}
	}
	return f[m][n]
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

func main() {

}
