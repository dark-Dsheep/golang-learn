package main

/*
*
279. 完全平方数
*/
func numSquares(n int) int {
	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = 0x3f3f3f
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], min(f[i-1]+1, f[i-j*j]+1))
		}
	}
	return f[n]
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
