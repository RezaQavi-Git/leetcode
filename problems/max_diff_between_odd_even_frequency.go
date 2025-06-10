package problems

func MaxDifference(s string) int {
    mapping := [26]int{}
    for i := 0; i < len(s); i++ {
        mapping[s[i]-'a']++
    }

    maxOdd := 0
    minEven := len(s)

    for _, freq := range mapping {
        if freq == 0 {
            continue
        }
        if freq%2 == 1 {
            if freq > maxOdd {
                maxOdd = freq
            }
        } else {
            if freq < minEven {
                minEven = freq
            }
        }
    }

    return maxOdd - minEven
}
