package main

/**
300. 最长递增子序列
*/

func lengthOfLIS(nums []int) int {
	ans := 1
	f := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}

func main() {

}
