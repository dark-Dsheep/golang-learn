package main

//739. 每日温度

func dailyTemperatures(t []int) []int {
	n := len(t)
	st := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && t[st[0]] <= t[i] {
			st = st[1:]
		}
		if len(st) > 0 {
			ans[i] = st[0] - i
		}
		st = append([]int{i}, st...)
	}
	return ans
}

func main() {

}
