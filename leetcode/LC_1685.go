package main

/*
*1685. 有序数组中差绝对值之和
 */
func getSumAbsoluteDifferences(nums []int) (ans []int) {
	n := len(nums)
	// ans[i]表示nums[i]与数组其他元素的差值之和
	// 等价于以i为分界点, 求Sum(i左边元素的元素差值之和)+Sum(i右边元素的元素差值之和)
	// left = i左边的前缀和 - 需要减去i个nums[i]
	// right = i右边的前缀和 - 需要减去n-i个nums[i]
	sum := make([]int, n+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	for i, x := range nums {
		left := abs(sum[i] - i*x)
		right := abs(sum[n] - sum[i] - (n-i)*x)
		ans = append(ans, left+right)
	}
	return ans
}

//	func abs(x int) int {
//		if x < 0 {
//			return -x
//		}
//		return x
//	}
func main() {

}
