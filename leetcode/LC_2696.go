package main

//2696. 删除子串后的字符串最小长度

func minLength(s string) (ans int) {
	st := []byte{}
	for i := range s {
		v := s[i]
		if len(st) > 0 && (v == 'B' && st[len(st)-1] == 'A' || v == 'D' && st[len(st)-1] == 'C') {
			st = st[:len(st)-1]
		} else {
			st = append(st, v)
		}
	}
	ans = len(st)
	return
}

func main() {

}
