package main

//39. 组合总和

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(int, int)
	dfs = func(idx int, target int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		if target < 0 {
			return
		}
		for i := idx; i < len(candidates); i++ {
			comb = append(comb, candidates[i])
			dfs(i, target-candidates[i])
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, target)
	return ans
}

func main() {

}
