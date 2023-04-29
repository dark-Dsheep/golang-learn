package main

/*
*
198. 打家劫舍
*/
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	f := make([]int, 110)
	f[0], f[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		f[i] = max(f[i-1], f[i-2]+nums[i])
	}
	return f[n-1]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
