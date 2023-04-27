package main

/*
*
53. 最大子数组和
*/
func maxSubArray(nums []int) (ans int) {
	//n := len(nums)
	//f := make([]int, 100001)
	//f[0], ans = nums[0], nums[0]
	//for i := 1; i < n; i++ {
	//	f[i] = max(f[i-1]+nums[i], nums[i])
	//	ans = max(ans, f[i])
	//}
	//return ans
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i+1] > nums[i] {
			nums[i] += nums[i-1]
		}
		ans = max(ans, nums[i])
	}
	return ans
}

func main() {
}
