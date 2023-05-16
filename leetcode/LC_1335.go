package main

// 第I类区间DP

// 1335. 工作计划的最低难度
func minDifficulty(job []int, d int) int {
	n := len(job)
	if n < d {
		return -1
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, d+1)
	}
	job = append([]int{0}, job...)
	// 当d=1的时候 f[i][1]的值
	for i := 1; i <= n; i++ {
		// f[i][1] = calc(1, i)
		f[i][1] = max(f[i-1][1], job[i])
	}
	for i := 1; i <= n; i++ {
		for k := 2; k <= min(d, i); k++ {
			f[i][k] = 0x3f3f3f
			mx := 0
			for j := i; j >= k; j-- {
				mx = max(mx, job[j])
				f[i][k] = min(f[i][k], f[j-1][k-1]+mx)
			}
		}
	}
	return f[n][d]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
