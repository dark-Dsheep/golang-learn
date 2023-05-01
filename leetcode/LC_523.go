package main

/*
*523. 连续的子数组和
 */
func checkSubarraySum(nums []int, k int) bool {
	mp := make(map[int]int)
	mp[0] = -1
	sum := 0
	for i, x := range nums {
		sum += x
		if j, ok := mp[sum%k]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			mp[sum%k] = i
		}
	}
	return false
}

func main() {

}
