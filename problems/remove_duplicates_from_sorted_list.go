package problems

func listPrint(l []int) {
	for _, n := range l {
		print(n)
	}
	println()
}

func RemoveDuplicates(nums []int) int {
	for i := 0; i <= len(nums)-2; i++ {
		for nums[i+1] == nums[i] {
			if len(nums) == 2 {
				if nums[0] == nums[1] {
					return 1
				} else {
					return 2
				}

			}
			nums = append(nums[:i+1], nums[i+2:]...)
		}
	}

	return len(nums)
}
