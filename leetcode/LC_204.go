package main

/*
204. 计数质数
*/

func countPrimes(n int) (ans int) {
	st := make([]bool, n)
	for i := 2; i < n; i++ {
		if st[i] {
			continue
		}
		ans++
		for j := 2 * i; j < n; j += i {
			st[j] = true
		}
	}
	return ans
}

func main() {

}
