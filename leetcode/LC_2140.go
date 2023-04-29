package main

/*
*
2140. 解决智力问题
*/
func mostPoints(q [][]int) int64 {
	n := len(q)
	f := make([]int, n+1)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) (ans int) {
		if i >= len(q) {
			return 0
		}
		ptr := &f[i]
		if *ptr != -1 {
			return *ptr
		}
		ans = max(dfs(i+q[i][1]+1)+q[i][0], dfs(i+1))
		*ptr = ans
		return ans
	}
	return int64(dfs(0))
}

func main() {

}
