package problems

import (
	"strconv"
	"strings"
)


func RLE(s string) string {

	var sb strings.Builder
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		cnt := 1
		for i < lenS-1 && s[i] == s[i+1] {
			cnt++
			i++
		}
		sb.WriteString(strconv.Itoa(cnt))
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func CountAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	return RLE(CountAndSay(n-1))
}