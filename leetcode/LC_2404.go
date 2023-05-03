package main

/**
2404. 出现最频繁的偶数元素
*/

func mostFrequentEven(nums []int) (ans int) {
	mp := make(map[int]int)
	ans = 0x3f3f3f3f
	max := 0
	for _, x := range nums {
		if (x & 1) == 0 {
			mp[x]++
			if max < mp[x] {
				max = mp[x]
				ans = x
			} else if max == mp[x] {
				ans = min(ans, x)
			}
		}
	}
	if ans == 0x3f3f3f3f {
		return -1
	}
	return ans
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
