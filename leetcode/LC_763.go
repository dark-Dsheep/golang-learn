package main

/*
*
763. 划分字母区间
*/
func partitionLabels(s string) (ans []int) {
	last := make([]int, 126) // last[i]表示第i个字母最后出现的位置
	for i, x := range s {
		last[x] = i
	}
	start, end := 0, 0
	for i, x := range s {
		end = max(end, last[x])
		if end == i { // 将s[start:end]划为一个片段
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
