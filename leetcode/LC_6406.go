package main

/*
*
6406. K 个元素的最大和
*/
func maximizeSum(nums []int, k int) (ans int) {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	for k > 0 {
		ans += max
		max += 1
		k--
	}
	return ans
}

func main() {

}
