package main

/*
*
503. 下一个更大元素 II
*/
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	st := []int{}
	ans := make([]int, n)
	for i := 2*n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[st[0]] <= nums[i%n] {
			st = st[1:]
		}
		if len(st) > 0 {
			ans[i%n] = nums[st[0]]
		} else {
			ans[i%n] = -1
		}
		st = append([]int{i % n}, st...)
	}
	return ans
}

func main() {

}
