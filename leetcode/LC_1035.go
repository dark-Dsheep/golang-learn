package main

// 1035. 不相交的线
func maxUncrossedLines(A []int, B []int) int {
	// LCS问题
	m, n := len(A), len(B)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if A[i] == B[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[m][n]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
