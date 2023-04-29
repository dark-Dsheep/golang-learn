package main

/*
*
121. 买卖股票的最佳时机
*/
func maxProfit(prices []int) int {
	// f[0]表示不持有股票的最大利润
	// f[1]表示持有股票的最大利润
	f := make([]int, 2)
	f[1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		f[0] = max(f[0], f[1]+prices[i])
		f[1] = max(f[1], -prices[i])
	}
	return f[0]
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {

}
