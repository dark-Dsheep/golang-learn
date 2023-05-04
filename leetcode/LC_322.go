package main

/*
*
322. 零钱兑换
*/
func coinChange(coins []int, amount int) (ans int) {
	f := make([]int, 10001)
	var dfs func(int) int
	dfs = func(n int) (ans int) {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return 0x3f3f3f
		}
		ptr := &f[n]
		if *ptr != 0 {
			return *ptr
		}
		ans = 0x3f3f3f
		for _, x := range coins {
			sub := dfs(n - x)
			if sub != 0x3f3f3f {
				ans = min(ans, sub+1)
			}
		}
		*ptr = ans
		return ans
	}
	ans = dfs(amount)
	if ans == 0x3f3f3f {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
