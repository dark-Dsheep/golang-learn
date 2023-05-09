package main

/*
*
2437. 有效时间的数目
*/
func countTime(time string) int {
	h, m := 0, 0
	if time[0] == '?' {
		if time[1] == '?' {
			h = 24
		} else if time[1] >= '4' {
			h = 2
		} else {
			h = 3
		}
	} else if time[0] == '2' {
		if time[1] == '?' {
			h = 4
		} else {
			h = 1
		}
	} else {
		if time[1] == '?' {
			h = 10
		} else {
			h = 1
		}
	}
	if time[3] == '?' {
		if time[4] == '?' {
			m = 60
		} else {
			m = 6
		}
	} else {
		if time[4] == '?' {
			m = 10
		} else {
			m = 1
		}
	}
	return h * m
}

func main() {

}
