package main

/*
*1043. 分隔数组以得到最大和
 */
func maxSumAfterPartitioning(arr []int, k int) int {
	f := make([]int, 510)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) (ans int) {
		if i < 0 {
			return 0
		}
		ptr := &f[i]
		if *ptr != -1 {
			return *ptr
		}
		mx := 0
		for j := i; i-j+1 <= k && j >= 0; j-- {
			mx = max(mx, arr[j])
			ans = max(ans, dfs(j-1)+mx*(i-j+1))
		}
		*ptr = ans
		return ans
	}
	return dfs(len(arr) - 1)
}

func main() {

}
