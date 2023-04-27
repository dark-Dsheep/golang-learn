package main

/**
2597. 美丽子集的数目
*/

func beautifulSubsets(nums []int, k int) int {
	mp := make([]int, 1001+k*2)
	ans := -1
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			ans += 1
			return
		}
		dfs(i + 1)
		x := nums[i] + k
		if mp[x-k] == 0 && mp[x+k] == 0 {
			mp[x]++
			dfs(i + 1)
			mp[x]--
		}
	}
	dfs(0)
	return ans
}

func main() {

}
