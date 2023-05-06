package main

/*
*1419. 数青蛙
 */
var previous = [...]int{'c': 'k', 'r': 'c', 'o': 'r', 'a': 'o', 'k': 'a'}

func minNumberOfFrogs(croakOfFrogs string) int {
	cnt := [len(previous)]int{}
	for _, ch := range croakOfFrogs {
		pre := previous[ch]
		if cnt[pre] > 0 {
			cnt[pre]--
		} else if ch != 'c' {
			return -1
		}
		cnt[ch]++
	}
	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1
	}
	return cnt['k']
}

func main() {

}
