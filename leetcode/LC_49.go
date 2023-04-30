package main

import "sort"

/*
*
49. 字母异位词分组
*/
func groupAnagrams(strs []string) (ans [][]string) {
	mp := make(map[string][]string)
	for _, str := range strs {
		cs := []byte(str)
		sort.Slice(cs, func(i, j int) bool {
			return cs[i] < cs[j]
		})
		s := string(cs)
		mp[s] = append(mp[s], str)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func main() {

}
