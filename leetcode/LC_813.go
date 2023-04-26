package main

//813. 最大平均值和的分组

func largestSumOfAverages(nums []int, k int) float64 {
	N := len(nums)
	K := k
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = nums[i-1]
	}
	var cal func(int, int) float64
	cal = func(j int, i int) float64 {
		cnt := float64(i - j + 1)
		sum := 0.00
		for ; j <= i; j++ {
			sum += float64(a[j])
		}
		return sum / cnt
	}
	f := make([][110]float64, 110)
	for i := 1; i <= N; i++ {
		f[i][1] = cal(1, i)
	}
	for i := 1; i <= N; i++ {
		for k = 2; k <= min(i, K); k++ {
			for j := i; j >= k; j-- {
				f[i][k] = Max(f[i][k], f[j-1][k-1]+cal(j, i))
			}
		}
	}
	return f[N][K]
}

func Max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
