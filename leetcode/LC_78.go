package main

/*
*
78. 子集
*/
func subsets(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(start int) {
		if start == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[start])
		dfs(start + 1)
		set = set[:len(set)-1]
		dfs(start + 1)
	}
	dfs(0)
	return ans
}

func main() {

}
