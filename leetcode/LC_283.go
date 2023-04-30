package main

/*
*
283. 移动零
*/
func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

func main() {

}
