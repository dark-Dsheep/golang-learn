package main

/*
*
239. 滑动窗口最大值
*/
func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := []int{}
	for i, x := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1] // 移除队列中比新元素小的队尾元素
		}
		q = append(q, i)
		if i >= k-1 {
			for len(q) > 0 && q[0] < i-k+1 {
				q = q[1:] // 移除不在当前窗口的队头下标
			}
			ans = append(ans, nums[q[0]])
		}
	}
	return ans
}

func main() {

}
