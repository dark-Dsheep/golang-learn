package main

/*
*
560. 和为 K 的子数组
*/
func subarraySum(nums []int, k int) (ans int) {
	mp := make(map[int]int)
	mp[0] = 1
	sum := 0
	for _, x := range nums {
		sum += x
		if cnt, ok := mp[sum-k]; ok {
			ans += cnt
		}
		mp[sum]++
	}
	return ans
}

func main() {

}
