package main

import "sort"

/*
*
6367. 矩阵中的和
*/
func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}
	m, n := len(nums), len(nums[0])
	for j := 0; j < n; j++ {
		score := 0
		for i := 0; i < m; i++ {
			score = max(ans, nums[i][j])
		}
		ans += score
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
