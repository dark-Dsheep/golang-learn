package main

/**
91. 解码方法
*/

func numDecodings(s string) int {
	f := make([]int, 110)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) (ans int) {
		ptr := &f[i]
		if i >= len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		if *ptr != -1 {
			return *ptr
		}
		ans = dfs(i + 1) // 单独解码
		if i+2 <= len(s) && ((s[i]-'0')*10+(s[i+1]-'0')) <= 26 {
			ans += dfs(i + 2) // 与后一个字符组合解码
		}
		*ptr = ans
		return ans
	}
	return dfs(0)
}

func main() {
	//str := "226"
	//x := (str[0]-'0')*10 + (str[1] - '0')
	//fmt.Println(x)
}
