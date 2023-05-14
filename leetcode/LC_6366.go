package main

/*
*
6367. 矩阵中的和
*/
func countSeniors(details []string) (ans int) {
	for _, d := range details {
		// &15 等价于 -'0'
		if d[11]&15*10+d[12]&15 > 60 {
			ans++
		}
	}
	return ans
}

func main() {

}
