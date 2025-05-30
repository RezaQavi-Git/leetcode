package problems

func convertToMorse(s string) string {

	morseCodeMap := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	c := ""

	for i := 0; i < len(s); i++ {
		c += morseCodeMap[s[i]]
	}

	return c
}

func UniqueMorseRepresentations(words []string) int {
	set := make(map[string]bool)
	wordsSet := make(map[string]string)

	for _, w := range words {
		mw, wfound := wordsSet[w]
		m := ""
		if !wfound {
			m = convertToMorse(w)
			wordsSet[w] = m
		} else {
			m = mw
		}
				
		_, found := set[m]
		if !found {
			set[m] = true
		}
	}

	return len(set)

}
