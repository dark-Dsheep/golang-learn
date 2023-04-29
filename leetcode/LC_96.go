package main

/**
96. 不同的二叉搜索树
*/

func numTrees(n int) int {
	f := make([]int, 20)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(x int) (ans int) {
		if x == 0 || x == 1 {
			return 1
		}
		ptr := &f[x]
		if *ptr != -1 {
			return *ptr
		}
		for i := 1; i <= x; i++ {
			ans += dfs(i-1) * dfs(x-i)
		}
		*ptr = ans
		return ans
	}
	return dfs(n)
}

func main() {

}
