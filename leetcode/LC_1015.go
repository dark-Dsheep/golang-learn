package main

/*
*
1015. 可被 K 整除的最小整数
*/
func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	x := 1 % k
	for i := 1; ; i++ {
		if x == 0 {
			return i
		}
		x = (x*10 + 1) % k
	}
}

func main() {

}
