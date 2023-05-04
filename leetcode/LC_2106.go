package main

/**
2106. 摘水果
*/

func maxTotalFruits(fruits [][]int, startPos, k int) int {
	f := make([]int, 200001)
	for _, v := range fruits {
		f[v[0]] += v[1]
	}
	for i := 1; i <= 200000; i++ {
		f[i] += f[i-1]
	}
	ans := 0
	// 两种走法: 1.往右走x, 折返往左走y 2.往左走x, 折返往右走y
	// 两种走法都走了, 2 * x + y步 => k = 2 * x + y 推出 y = k - 2 * x
	// 枚举x
	for x := 0; 2*x <= k; x++ {
		y := k - 2*x
		l1 := max(startPos-x, 0)
		r1 := min(startPos+y, 200000)
		l2 := max(startPos-y, 0)
		r2 := min(startPos+x, 200000)
		v1, v2 := 0, 0
		if l1 > 0 {
			v1 = f[l1-1]
		}
		if l2 > 0 {
			v2 = f[l2-1]
		}
		ans = max(ans, max(f[r1]-v1, f[r2]-v2))
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {
	fruits := [][]int{{0, 10000}}
	maxTotalFruits(fruits, 200000, 200000)
}
