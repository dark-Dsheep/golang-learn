package main

/**
376. 摆动序列
*/

func wiggleMaxLength(nums []int) int {
	f := make([]int, 2)
	f[0], f[1] = 1, 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] < nums[i-1] {
			f[1] = f[0] + 1
		} else if nums[i] > nums[i-1] {
			f[0] = f[1] + 1
		}
		ans = max(ans, max(f[0], f[1]))
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
