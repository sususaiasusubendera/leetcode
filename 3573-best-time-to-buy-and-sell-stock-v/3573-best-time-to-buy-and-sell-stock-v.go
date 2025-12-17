func maximumProfit(prices []int, k int) int64 {
	dp := make([][3]int64, len(prices))
	// day 0
	for i := 1; i <= k; i++ {
        // OPEN
		dp[i][1] = int64(-prices[0]) // normal: buy -> sell
		dp[i][2] = int64(prices[0])  // short: sell -> buy
	}

	for i := 1; i < len(prices); i++ {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], max(dp[j][1]+int64(prices[i]), dp[j][2]-int64(prices[i]))) // idle, CLOSE
            dp[j][1] = max(dp[j][1], dp[j-1][0]-int64(prices[i])) // normal, OPEN
            dp[j][2] = max(dp[j][2], dp[j-1][0]+int64(prices[i])) // short, OPEN
		}
	}

	return dp[k][0]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// array, dp (bot-up rolling state)
// time: O(nk)
// space: O(k)