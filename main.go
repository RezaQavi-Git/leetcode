package main
import "leetcode/problems"

func main() {
	s := []int{1, 1, 2, 3, 4, 4, 5, 5, 5}
	println(problems.RemoveDuplicates(s))
	for _, n := range(s) {
		print(n)
	}
	println()
}