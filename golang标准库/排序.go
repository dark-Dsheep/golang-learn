package main

import (
	"fmt"
	"sort"
)

type Person struct {
	UserId   string
	Username string
	Age      int
}

type PersonSlice []Person

// 返回切片长度
func (p PersonSlice) Len() int {
	return len(p)
}

// compare
func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// switch
func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// 整型排序
	arr := []int{1, 3, 2, 5, 7, 4}
	sort.Ints(arr)
	fmt.Println(arr)

	// 浮点排序
	floats := []float64{1.0, 2.4, 3.8, 1.11, 5.5}
	sort.Float64s(floats)
	fmt.Println(floats)

	// 字符串排序
	strings := []string{"helloworld", "aaa", "bbb", "ccc"}
	sort.Strings(strings)
	fmt.Println(strings)

	// 逆向排序
	nums := []int{1, 2, 4, 5, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	// 自定义排序
}
