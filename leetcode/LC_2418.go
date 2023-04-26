package main

import "sort"

/**
2418. 按身高排序
*/

func sortPeople(names []string, heights []int) []string {
	//mp := make(map[int]string)
	//for i, h := range heights {
	//	mp[h] = names[i]
	//}
	//sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	//ans := make([]string, len(names))
	//for i, h := range heights {
	//	ans[i] = mp[h]
	//}
	//return ans
	n := len(names)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return heights[id[i]] > heights[id[j]]
	})
	ans := make([]string, n)
	for i, j := range id {
		ans[i] = names[j]
	}
	return ans
}

func main() {

}
