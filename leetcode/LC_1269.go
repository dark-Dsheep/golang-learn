package main

/*
*1269. 停在原地的方案数
 */

func numWays(steps int, arrLen int) int {
	f := make([][510]int, 510)
	for i, row := range f {
		for j := range row {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(pos int, steps int) (ans int) {
		if pos < 0 || pos >= arrLen {
			return 0
		}
		if steps == 0 {
			if pos == 0 {
				return 1
			} else {
				return 0
			}
		}
		ptr := &f[pos][steps]
		if *ptr != -1 {
			return *ptr
		}
		ans = (ans + dfs(pos-1, steps-1)) % 1000000007
		ans = (ans + dfs(pos+1, steps-1)) % 1000000007
		ans = (ans + dfs(pos, steps-1)) % 1000000007
		*ptr = ans
		return ans
	}
	return dfs(0, steps)
}

func main() {

}
