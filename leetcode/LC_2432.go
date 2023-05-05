package main

/*
*
2432. 处理用时最长的那个任务的员工
*/
func hardestWorker(n int, logs [][]int) (ans int) {
	ans = logs[0][0]
	max := logs[0][1]
	for i := 1; i < len(logs); i++ {
		v := logs[i][1] - logs[i-1][1]
		if v > max {
			max = v
			ans = logs[i][0]
		} else if v == max {
			ans = min(ans, logs[i][0])
		}
	}
	return ans
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func main() {

}
