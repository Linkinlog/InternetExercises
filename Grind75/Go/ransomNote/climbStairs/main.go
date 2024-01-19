package main

func climbStairs(n int) int {
	// Create a 1D dp where dp[i] represent the number of ways to reach the ith stair from the bottom.
	// Initialise dp[0] = 1, as there is only one way for n = 0 and dp[1] = 2 as there are only 2 ways for input n = 2.
	// Now for each i >= 2, dp[i] = dp[i-1]+dp[i-2]

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]+dp[i-2]
	}
	return dp[n]
}
