func maximumProfit(prices []int, k int) int64 {
	dp := make([][][]int64, len(prices))
	for i := range dp {
		dp[i] = make([][]int64, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int64, 3)
		}
	}

    // day 0
    for j := 1; j <= k; j++ {
        dp[0][j][1] = int64(-prices[0]) // normal: buy -> sell
        dp[0][j][2] = int64(prices[0]) // short: sell -> buy
    }

    for i := 1; i < len(prices); i++ {
        for j := 1; j <= k; j++ {
            dp[i][j][0] = max(dp[i-1][j][0], max(dp[i-1][j][1]+int64(prices[i]), dp[i-1][j][2]-int64(prices[i])))
            dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-int64(prices[i]))
            dp[i][j][2] = max(dp[i-1][j][2], dp[i-1][j-1][0]+int64(prices[i]))
        }
    }

    return dp[len(prices)-1][k][0]
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// array, dp (bot-up)
// time: O(nk)
// space: O(nk)
