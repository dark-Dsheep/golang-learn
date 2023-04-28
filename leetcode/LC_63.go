package main

/*
*
63. 不同路径 II
*/
func uniquePathsWithObstacles(a [][]int) int {
	m, n := len(a), len(a[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i, row := range a {
		for j := range row {
			if a[i][j] == 1 {
				f[i][j] = 0
			} else {
				if i == 0 && j == 0 {
					f[i][j] = 1
				} else if i == 0 {
					f[i][j] = f[i][j-1]
				} else if j == 0 {
					f[i][j] = f[i-1][j]
				} else {
					f[i][j] = f[i-1][j] + f[i][j-1]
				}
			}
		}
	}
	return f[m-1][n-1]
}

func main() {

}
