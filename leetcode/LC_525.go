package main

/*
*
525. 连续数组
*/
func findMaxLength(nums []int) (ans int) {
	mp := make(map[int]int)
	mp[0] = -1
	sum := 0
	for i, x := range nums {
		if x == 1 {
			sum += 1
		} else {
			sum -= 1
		}
		if j, ok := mp[sum]; ok {
			ans = max(ans, i-j)
		} else {
			mp[sum] = i
		}
	}
	return ans
}

func main() {

}
