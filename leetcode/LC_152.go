package main

// 152. 乘积最大子数组
func maxProduct(nums []int) (ans int) {
	n := len(nums)
	// f[i][0] 表示以i结尾乘积为正数的最大子数组
	// f[i][1] 表示以i结尾乘积为负数的最大子数组
	f := make([][2]int, n)
	f[0][1], f[0][1], ans = nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0]*nums[i], max(f[i-1][1]*nums[i], nums[i]))
		f[i][1] = min(f[i-1][1]*nums[i], min(f[i-1][0]*nums[i], nums[i]))
		ans = max(ans, f[i][0])
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

}
