package main

import "sort"

/*
*
15. 三数之和
*/
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	// 固定i，枚举j,k => nums[i] + nums[j] + nums[k] == 0
	// 转换为 nums[j] + nums[k] = -nums[i]
	for i := range nums[:n-2] {
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复元素
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ { // 跳过重复数字
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- { // 跳过重复数字
				}
			}
		}
	}
	return ans
}

func main() {

}
