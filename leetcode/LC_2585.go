package main

/*
*
2585. 获得分数的方法数
*/
const mod = 1e9 + 7

func waysToReachTarget(target int, types [][]int) int {
	f := make([][1010]int, 60)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(idx int, target int) (ans int) {
		if target == 0 {
			return 1
		}
		if idx >= len(types) || target < 0 {
			return 0
		}
		ptr := &f[idx][target]
		if *ptr != -1 {
			return *ptr
		}
		ans = 0
		cnt := types[idx][0]
		for i := 0; i <= cnt; i++ {
			ans = (ans + dfs(idx+1, target-i*types[idx][1])) % mod
		}
		f[idx][target] = ans
		return ans
	}
	return dfs(0, target) % mod

}

func main() {

}
