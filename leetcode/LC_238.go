package main

/*
*
238. 除自身以外数组的乘积
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	left, right := 1, 1 // 左右累乘
	for i := range ans {
		ans[i] = 1
	}
	for i := range nums {
		ans[i] *= left
		left *= nums[i]
		ans[n-1-i] *= right
		right *= nums[n-1-i]
	}
	return ans
}

func main() {

}
