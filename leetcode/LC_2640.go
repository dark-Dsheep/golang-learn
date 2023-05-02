package main

/*
*
2640. 一个数组所有前缀的分数
*/
func findPrefixScore(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)
	ans[0] = int64(2 * nums[0])
	mx := nums[0]
	for i := 1; i < n; i++ {
		mx = max(mx, nums[i])
		ans[i] = ans[i-1] + int64(mx+nums[i])
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
