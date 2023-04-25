package main

/**
2653. 滑动子数组的美丽值
*/

func getSubarrayBeauty(nums []int, k int, x int) []int {
	const offset = 50
	cnt := [offset*2 + 1]int{}
	for _, num := range nums[:k-1] {
		cnt[num+offset]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+offset]++
		left := x
		for j, c := range cnt[:offset] {
			left -= c
			if left <= 0 {
				ans[i] = j - offset
				break
			}
		}
		cnt[nums[i]+offset]--
	}
	return ans
}

func main() {

}
