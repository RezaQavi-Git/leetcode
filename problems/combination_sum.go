package problems

func CombinationSum(candidates []int, target int) [][]int {
    var res [][]int
	var comb []int

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
		if target == 0 {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}
			comb = append(comb, candidates[i])
			backtrack(i, target - candidates[i])
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0, target)
	return res
}