package problems

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	rob1 := nums[0]
	rob2 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		current := max(nums[i]+rob1, rob2)
		rob1, rob2 = rob2, current
	}

	return rob2
}
