package main

/*
*
992. K 个不同整数的子数组
*/
func subarraysWithKDistinct(nums []int, k int) (ans int) {
	var calc func(int) int
	calc = func(k int) (ret int) {
		mp := make([]int, 20001)
		j, cnt := 0, 0
		for i, x := range nums {
			if mp[x] == 0 {
				cnt++
			}
			mp[x]++
			for cnt > k {
				mp[nums[j]]--
				if mp[nums[j]] == 0 {
					cnt -= 1
				}
				j++
			}
			ret += i - j + 1
		}
		return ret
	}
	return calc(k) - calc(k-1)
}

func main() {

}
