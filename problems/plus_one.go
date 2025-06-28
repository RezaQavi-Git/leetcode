package problems

func PlusOne(digits []int) []int {
	l := len(digits)
	c := 1
	for i := l - 1; i >= 0; i-- {
		digits[i] = digits[i] + c
		if digits[i] >= 10 {
			digits[i] -= 10
		} else {
			c = 0
			break
		}
	}

	if c == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
