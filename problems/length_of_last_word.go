package problems


func LengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	length := 0
	numberOfWrods := 0
	for i := len(s)-1; i >= 0; i-- {
		if numberOfWrods == 0 && s[i] == ' ' {
			continue
		} else if numberOfWrods == 1 && s[i] == ' ' {
			return length
		} else if numberOfWrods == 1 && s[i] != ' ' {
			length++
		} else {
			numberOfWrods++
			length++
		}
	}

	return length
}