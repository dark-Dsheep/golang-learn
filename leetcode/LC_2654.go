package main

import "math"

/*
*
2654. 使数组所有元素变成 1 的最少操作次数
*/
func minOperations(nums []int) int {
	n := len(nums)
	cnt := 0 // 记录nums中1的个数
	for _, x := range nums {
		if x == 1 {
			cnt++
		}
	}
	// 因为gcd(x,1)=1, 如果nums中有1,那么将所有元素变为1的最少操作次数为n-cnt
	if cnt > 0 {
		return n - cnt
	}
	// 固定区间左端点，枚举区间右端点，最小化一个子数组gcd=1的操作次数
	ans := math.MaxInt
	for i, x := range nums {
		sum := 0 // gcd操作次数
		for j := i + 1; j < n; j++ {
			x = gcd(x, nums[j])
			sum++
			if x == 1 {
				ans = min(ans, sum)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return n + ans - 1 // 最后还需要将剩下的n-1个元素变为1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
