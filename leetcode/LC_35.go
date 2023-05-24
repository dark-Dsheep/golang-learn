package main

// 35. 搜索插入位置
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {

}
