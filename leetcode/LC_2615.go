package main

/*
*
2615. 等值距离和
*/
func distance(nums []int) []int64 {
	g := map[int][]int{}
	for i, x := range nums {
		g[x] = append(g[x], i) // 将元素值相同的下标分组
	}
	ans := make([]int64, len(nums))
	for _, a := range g {
		n := len(a)
		sum := make([]int, n+1) // 等值下标距离前缀和
		for i, x := range a {
			sum[i+1] = sum[i] + x
		}
		for i, idx := range a {
			left := idx*i - sum[i]               // 左边与idx等值的距离和
			right := sum[n] - sum[i] - (n-i)*idx // 右边与idx等值的距离和
			ans[idx] = int64(left + right)
		}
	}
	return ans
}

func main() {

}
