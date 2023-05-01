package main

/*
*
496. 下一个更大元素 I
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := []int{}
	mp := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(st) > 0 && st[0] <= nums2[i] {
			st = st[1:]
		}
		if len(st) > 0 {
			mp[nums2[i]] = st[0]
		} else {
			mp[nums2[i]] = -1
		}
		st = append([]int{nums2[i]}, st...) // 头部插入元素 => 入栈
	}
	for i, x := range nums1 {
		nums1[i] = mp[x]
	}
	return nums1
}

func main() {

}
