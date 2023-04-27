package main

/*
*
22. 括号生成
*/
func generateParenthesis(n int) []string {
	var ans []string
	var dfs func(int, int, string)
	dfs = func(lcnt int, rcnt int, sb string) {
		if lcnt < 0 || rcnt < 0 {
			return
		}
		if lcnt == 0 && rcnt == 0 {
			ans = append(ans, sb)
			return
		}
		if lcnt > rcnt {
			return
		}
		dfs(lcnt-1, rcnt, sb+"(")
		dfs(lcnt, rcnt-1, sb+")")
	}
	dfs(n, n, "")
	return ans
}

func main() {

}
