package problems

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func EditDistance(w1, w2 string) int {
	m := len(w1)
	n := len(w2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insert := dp[i][j-1] + 1
				delete := dp[i-1][j] + 1
				replace := dp[i-1][j-1] + 1
				dp[i][j] = min(insert, min(delete, replace))
			}
		}
	}

	return dp[m][n]
}