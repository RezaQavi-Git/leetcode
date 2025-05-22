package problems


func ConvertToInteger(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	num := mapping[s[l-1]]
	for i := l-2; i >= 0; i-- {
		if mapping[s[i]] >= mapping[s[i+1]] {		
			num += mapping[s[i]]
		} else {
			num -= mapping[s[i]]
		}
	}
	return num
}
