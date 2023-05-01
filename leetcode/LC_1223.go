package main

/**
1223. 掷骰子模拟
*/

func dieSimulator(n int, rollMax []int) int {
	f := make([][16][5010]int, 7)
	for i, row := range f {
		for j, col := range row {
			for k := range col {
				f[i][j][k] = -1
			}
		}
	}
	rollMax = append([]int{0}, rollMax...)
	var dfs func(int, int, int) int
	dfs = func(num int, cnt int, n int) (ans int) {
		if cnt > rollMax[num] {
			return 0
		}
		if n == 0 {
			return 1
		}
		ptr := &f[num][cnt][n]
		if *ptr != -1 {
			return *ptr
		}
		for i := 1; i <= 6; i++ {
			if num == i {
				ans = (ans + dfs(i, cnt+1, n-1)) % 1000000007
			} else {
				ans = (ans + dfs(i, 1, n-1)) % 1000000007
			}
		}
		*ptr = ans
		return ans
	}
	return dfs(0, 0, n)
}

func main() {

}
