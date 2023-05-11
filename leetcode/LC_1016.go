package main

/*
*1016. 子串能表示从 1 到 N 数字的二进制串
 */
func queryString(s string, n int) bool {
	// 暴力
	//for i := 1; i <= n; i++ {
	//	if !strings.Contains(s, strconv.FormatUint(uint64(i), 2)) {
	//		return false
	//	}
	//}
	//return true

	vis := map[int]struct{}{}
	for i, b := range s {
		x := int(b - '0')
		if x == 0 {
			continue
		}
		for j := i + 1; x <= n; i++ {
			vis[x] = struct{}{}
			if j == len(s) {
				break
			}
			x = x<<1 | int(s[j]-'0')
		}
	}
	return len(vis) == n
}

func main() {

}
