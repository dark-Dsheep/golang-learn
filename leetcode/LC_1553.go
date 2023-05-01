package main

/*
*
1553. 吃掉 N 个橘子的最少天数
*/
func minDays(n int) int {
	mp := make(map[int]int)
	var dfs func(int) int
	dfs = func(n int) (ans int) {
		if n < 3 {
			return n
		}
		if v, ok := mp[n]; ok {
			return v
		}
		x, y := n%2, n%3
		ans = min(x+dfs(n/2)+1, y+dfs(n/3)+1)
		mp[n] = ans
		return ans
	}
	return dfs(n)
}

func main() {

}
