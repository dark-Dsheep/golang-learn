package main

import "fmt"

/**
1. 两数之和
*/

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, val := range nums {
		v := target - val
		if j, f := mp[v]; f {
			return []int{j, i}
		}
		mp[val] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
