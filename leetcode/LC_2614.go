package main

/*
*
2614. 对角线上的质数
*/
func diagonalPrime(nums [][]int) (ans int) {
	n := len(nums)
	var check func(int) bool
	check = func(x int) bool {
		if x < 2 {
			return false
		}
		for i := 2; i <= x/i; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := range nums {
		if check(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
		if check(nums[i][n-i-1]) {
			ans = max(ans, nums[i][n-i-1])
		}
	}
	return ans
}

func main() {

}
