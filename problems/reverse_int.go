package problems

func Reverse(x int) int {
	r, negflag := 0, 1
	if x < 0 {
		negflag = -1
		x *= negflag
	}

	for x > 0 {
		r = r*10 + x%10
		x = x / 10
		if r > 2147483647 {
			return 0
		}
	}

	return r * negflag
}
