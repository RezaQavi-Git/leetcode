package main

import (
	"fmt"
	"leetcode/problems"
)

func main() {
	l1 := problems.ListNode{
		Val: 2, Next: &problems.ListNode{Val: 4, Next: &problems.ListNode{Val: 3}}}

	l2 := problems.ListNode{
		Val: 5, Next: &problems.ListNode{Val: 6, Next: &problems.ListNode{Val: 4}}}

		
	fmt.Println(
		"Sum: ",
		problems.AddTwoNumbers(&l1, &l2))
}
