package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 0, 0)
	nums = append(nums, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(nums), cap(nums))

	// 切片插入元素
	// 从头部插入
	nums = append([]int{-1, 0}, nums...)
	// 向中间下标i插入元素
	i := 1
	nums = append(nums[:i+1], append([]int{999, 999}, nums[i+1:]...)...)
	fmt.Println(nums)
	// 从尾部插入元素
	nums = append(nums, 99, 100)
	fmt.Println(nums)

	// 删除元素
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 从头部删除n个元素
	n := 3
	nums = nums[n:]
	fmt.Println(nums)
	// 从尾部删除n个元素
	nums = nums[:len(nums)-n]
	fmt.Println(nums)
	// 从指定下标i位置开始删除n个元素
	nums = append(nums[:i], nums[i+n:]...)
	fmt.Println(nums)
	// 删除所有元素
	nums = nums[:0]
	fmt.Println(nums)

	// 拷贝
	dest := make([]int, 0)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src))
	fmt.Println(src, dest)

	// 遍历
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range
	for idx, val := range slice {
		fmt.Println(idx, val)
	}

	// 多维切片

}
