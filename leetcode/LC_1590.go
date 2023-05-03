package main

/**
1590. 使数组和能被 P 整除
*/

func minSubarray(nums []int, p int) int {
	x, s := 0, 0
	ans := 0x3f3f3f
	for _, v := range nums {
		x = (x + v) % p
	}
	mp := make(map[int]int)
	mp[0] = -1
	for i, v := range nums {
		s = (s + v) % p
		mp[s] = i
		remain := ((s-x)%p + p) % p
		if j, ok := mp[remain]; ok {
			ans = min(ans, i-j)

		}
	}
	if ans < len(nums) {
		return ans
	}
	return -1
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}

func main() {

}
