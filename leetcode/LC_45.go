package main

/**
45. 跳跃游戏 II
*/

func jump(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := range f {
		f[i] = 0x3f3f3f
	}
	f[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if f[j] < 0x3f3f3f && j+nums[j] >= i {
				f[i] = f[j] + 1
				break
			}
		}
	}
	return f[n-1]
}

func main() {

}
