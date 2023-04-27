package main

/*
*
2606. 找到最大开销的子字符串
*/
func maximumCostSubstring(s string, chars string, vals []int) (ans int) {
	//mp := make(map[int]int, 130)
	//for i := 0; i < len(chars); i++ {
	//	mp[int(chars[i])] = vals[i]
	//}
	//n := len(s)
	//s = " " + s
	//f := make([]int, n+1)
	//f[0] = 0
	//for i := 1; i <= n; i++ {
	//	k := int(s[i])
	//	if val, b := mp[k]; b {
	//		f[i] = max(f[i-1]+val, val)
	//	} else {
	//		v := k - 'a' + 1
	//		f[i] = max(f[i-1]+v, v)
	//	}
	//	ans = max(ans, f[i])
	//}
	//return ans
	mp := [26]int{}
	for i := range mp {
		mp[i] = i + 1
	}
	for i, c := range chars {
		mp[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mp[c-'a']
		ans = max(ans, f)
	}
	return ans
}

func main() {

}
