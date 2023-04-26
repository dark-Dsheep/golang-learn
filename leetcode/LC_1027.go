package main

/*
*
1027. 最长等差数列
*/
func longestArithSeqLength(nums []int) int {
	f := make([][1010]int, 1010)
	for i := range f {
		for j := range f[i] {
			f[i][j] = 1
		}
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j] + 500
			f[i][d] = max(f[i][d], f[j][d]+1)
			ans = max(ans, f[i][d])
		}
	}
	return ans
}

func main() {

}
