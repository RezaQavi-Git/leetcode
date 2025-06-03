package main

import (
	"fmt"
	"leetcode/problems"
)

func main() {
	fmt.Println(
		"Min Distance:",
		problems.EditDistance("intention", "execution"))
}
