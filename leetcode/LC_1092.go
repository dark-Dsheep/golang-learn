package main

/**
1092. 最短公共超序列
*/

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	str1, str2 = " "+str1, " "+str2
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i] == str2[j] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	ans := []byte{}
	i, j := m, n
	for i > 0 && j > 0 {
		if str1[i] == str2[j] {
			ans = append(ans, str1[i])
			i--
			j--
		} else if f[i][j] == f[i-1][j] {
			ans = append(ans, str1[i])
			i--
		} else if f[i][j] == f[i][j-1] {
			ans = append(ans, str2[j])
			j--
		}
	}
	for i > 0 {
		ans = append(ans, str1[i])
		i--
	}
	for j > 0 {
		ans = append(ans, str2[j])
		j--
	}
	return reverse(ans)
}

func reverse(b []byte) string {
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {

}
