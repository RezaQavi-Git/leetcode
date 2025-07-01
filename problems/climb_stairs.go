package problems


func ClimbStairs(n int) int {
	
	dp := make(map[int]int)
	
	dp[1], dp[2] = 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}