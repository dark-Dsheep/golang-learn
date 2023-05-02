package main

import "sort"

/*
*56. 合并区间
 */
func merge(intervals [][]int) [][]int {
	list := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	list = append(list, intervals[0])
	for i := 1; i < len(intervals); i++ {
		end := &list[len(list)-1][1]
		if intervals[i][0] <= *end {
			*end = max(*end, intervals[i][1])
		} else {
			list = append(list, intervals[i])
		}
	}
	return list
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
