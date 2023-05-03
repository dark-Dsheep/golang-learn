package main

/*
*
1003. 检查替换后的词是否有效
*/
func isValid(s string) bool {
	st := []int{}
	for _, ch := range s {
		if ch == 'a' || ch == 'b' {
			st = append([]int{int(ch)}, st...)
		} else {
			cnt := 2
			for len(st) > 0 && cnt > 0 {
				if cnt == 2 && st[0] == 'b' {
					st = st[1:]
					cnt--
				} else if cnt == 1 && st[0] == 'a' {
					st = st[1:]
					cnt--
				} else {
					return false
				}
			}
			if cnt != 0 {
				return false
			}
		}
	}
	return len(st) == 0
}
