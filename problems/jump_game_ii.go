package problems

func Jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps := 0
	curEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == curEnd {
			jumps++
			curEnd = farthest
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
