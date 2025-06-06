package problems

func RobotWithString(s string) string {
	n := len(s)

	count := [26]int{}
	for i := 0; i < n; i++ {
		count[s[i]-'a']++
	}

	result := make([]byte, 0, n)
	stack := make([]byte, 0, n)

	minChar := byte('a')

	for i := 0; i < n; i++ {
		ch := s[i]
		count[ch-'a']--

		stack = append(stack, ch)

		for minChar <= 'z' && count[minChar-'a'] == 0 {
			minChar++
		}

		for len(stack) > 0 && stack[len(stack)-1] <= minChar {
			last := len(stack) - 1
			result = append(result, stack[last])
			stack = stack[:last]
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
	}

	return string(result)
}
