package main

import "fmt"

/*
*189. 轮转数组
 */
func rotate(nums []int, k int) {
	//n := len(nums)
	//arr := make([]int, n)
	//for i := range nums {
	//	j := (i + k) % n
	//	arr[j] = nums[i]
	//}
	//for i := range nums {
	//	nums[i] = arr[i]
	//}

	// O(1)空间
	var reverse func([]int)
	reverse = func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	//fmt.Println(nums)
	//fmt.Println(k)
	n := len(nums)
	fmt.Println((6 + k) % n)
}
