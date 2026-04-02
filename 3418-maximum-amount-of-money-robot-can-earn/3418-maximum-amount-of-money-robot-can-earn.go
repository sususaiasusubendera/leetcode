func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	const INF = int(1e18)

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
			for k := range dp[i][j] {
				dp[i][j][k] = -INF
			}
		}
	}

	// base
	if coins[0][0] >= 0 {
		dp[0][0][0] = coins[0][0]
		dp[0][0][1] = coins[0][0]
		dp[0][0][2] = coins[0][0]
	} else {
		dp[0][0][0] = coins[0][0]
		dp[0][0][1] = 0
		dp[0][0][2] = 0
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 { continue } // base case
            for k := 0; k <= 2; k++ {
                best := -INF
                
                // take from top
                if i > 0 {
                    // not using neutralize
                    best = max(best, dp[i-1][j][k] + coins[i][j])

                    // using neutralize
                    if coins[i][j] < 0 && k > 0 {
                        best = max(best, dp[i-1][j][k-1] + 0)
                    }
                }

                // take from left
                if j > 0 {
                    // not using neutralize
                    best = max(best, dp[i][j-1][k] + coins[i][j])

                    // using neutralize
                    if coins[i][j] < 0 && k > 0 {
                        best = max(best, dp[i][j-1][k-1] + 0)
                    }
                }

                dp[i][j][k] = best
            }
        }
    }

    return max(dp[m-1][n-1][0], dp[m-1][n-1][1], dp[m-1][n-1][2])
}

// array, dp (bot-up), matrix
// time: O(mn)
// space: O(mn)