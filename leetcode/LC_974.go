package main

/*
*974. 和可被 K 整除的子数组
 */
func subarraysDivByK(nums []int, k int) (ans int) {
	mp := make([]int, k)
	mp[0] = 1
	sum := 0
	for _, x := range nums {
		sum += x
		mod := (sum%k + k) % k
		ans += mp[mod]
		mp[mod]++
	}
	return ans
}

func main() {

}
