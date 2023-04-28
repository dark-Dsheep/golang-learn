package main

/**
55. 跳跃游戏
*/

func canJump(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] != 0 && j+nums[j] >= i {
				f[i] = f[j]
				break
			}
			if nums[i-1] != 0 {
				f[i] = f[i-1]
				break
			}
		}
	}
	return f[n-1]
}

func main() {

}
