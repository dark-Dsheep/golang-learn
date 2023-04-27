package main

/*
*
44. 通配符匹配
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	s, p = " "+s, " "+p
	f := make([][]bool, m+1)
	for i, row := range f {
		f[i] = make([]bool, n+1)
		for j := range row {
			f[i][j] = false
		}
	}
	f[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j] != '*' {
			break
		}
		f[0][j] = true
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j] == '*' {
				f[i][j] = f[i][j-1] || f[i-1][j]
			} else {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '?')
			}
		}
	}
	return f[m][n]
}

func main() {

}
