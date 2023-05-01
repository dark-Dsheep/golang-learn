package main

/*
*
930. 和相同的二元子数组
*/
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	mp := make(map[int]int)
	mp[0] = 1
	sum := 0
	for _, x := range nums {
		sum += x
		if cnt, ok := mp[sum-goal]; ok {
			ans += cnt
		}
		mp[sum]++
	}
	return ans
}

func main() {

}
