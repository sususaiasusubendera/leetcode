func numberOfPaths(grid [][]int, k int) int {
	const mod = 1_000_000_007

    m, n := len(grid), len(grid[0])
    dp := make([][][]int, m)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = make([]int, k)
        }
    }

    dp[0][0][grid[0][0]%k] = 1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            v := grid[i][j]
            if i > 0 {
                for r := 0; r < k; r++ {
                    newR := (r + v) % k
                    dp[i][j][newR] = (dp[i][j][newR] + dp[i-1][j][r]) % mod
                }
            }

            if j > 0 {
                for r := 0; r < k; r++ {
                    newR := (r + v) % k
                    dp[i][j][newR] = (dp[i][j][newR] + dp[i][j-1][r]) % mod
                }
            }
        }
    }

    return dp[m-1][n-1][0]
}

// array, dp, matrix
// time: O(mnk)
// space: O(mnk)