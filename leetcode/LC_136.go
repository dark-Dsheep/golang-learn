package main

/*
136. 只出现一次的数字
*/
func singleNumber(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return ans
}

func main() {

}
