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

	// 枚举s的子串(二进制数),保存在set中,如果size(set)==n说明s的子串符合条件
	vis := map[int]struct{}{}
	for i, b := range s {
		x := int(b - '0')
		if x == 0 { // 二进制从 1开始 s=0110等价于110
			continue
		}
		for j := i + 1; x <= n; j++ {
			vis[x] = struct{}{}
			if j == len(s) {
				break
			}
			x = x<<1 | int(s[j]-'0') // 子串[i,j]的二进制数
		}
	}
	return len(vis) == n
}

func main() {

}
