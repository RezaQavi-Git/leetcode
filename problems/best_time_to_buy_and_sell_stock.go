package problems

func MaxProfit(prices []int) int {
	maxP := 0
	maxProfits := make([]int, len(prices))
	maxProfits[0] = 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] + maxProfits[i-1] > 0 {
			maxProfits[i] = prices[i] - prices[i-1] + maxProfits[i-1]
			if maxProfits[i] > maxP {
				maxP = maxProfits[i]
			}
		} else {
			maxProfits[i] = 0
		}
	}

	return maxP
}
