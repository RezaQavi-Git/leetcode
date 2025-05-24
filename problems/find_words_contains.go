package problems


func isContains(w string, x byte) bool {
	for j := 0; j < len(w); j++ {
		if w[j] == x {
			return true
		}
	}

	return false
}

func FindWordsContaining(words []string, x byte) []int {

	var out []int 

	for i, w := range words {
		if isContains(w, x) {
			out = append(out, i)
		}
	}

	return out
}
